package routers

import (
	"bit_score/controllers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type SetupSessionRouters struct {
	server *gin.Engine
	db     *mongo.Database
}

func NewSessionRouters(server *gin.Engine, db *mongo.Database) *SetupSessionRouters {
	return &SetupSessionRouters{
		server: server,
		db:     db,
	}
}

func (s *SetupSessionRouters) setupRouters() {
	controller := controllers.NewSessionsController()

	sessions := s.server.Group("/sessions")
	{
		sessions.POST("/", controller.CreateSession)
	}
}

func (s *SetupSessionRouters) Routers() {
	s.setupRouters()
}

func SetupSessionRoutes(server *gin.Engine, db *mongo.Database) {
	sessionRouter := NewSessionRouters(server, db)
	sessionRouter.Routers()
}
