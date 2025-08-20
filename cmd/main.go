package main

import (
	"bit_score/db"
	"bit_score/router"
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	database, err := db.ConnectDb()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connected successfully")

	server := gin.Default()

	server.GET("/health", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		err := database.Client().Ping(ctx, nil)
		if err != nil {
			c.JSON(500, gin.H{
				"status":   "error",
				"database": "disconnected",
				"error":    err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"status":    "ok",
			"database":  "connected",
			"timestamp": time.Now().Format(time.RFC3339),
		})
	})

	router.SetupAllRoutes(server)

	srv := &http.Server{
		Addr:    ":8182",
		Handler: server,
	}

	go func() {
		log.Printf("server listen at %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	} else {
		log.Println("Server exiting gracefully")
	}
}
