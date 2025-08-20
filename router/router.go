package router

import (
	"bit_score/routers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type SetupRoutes struct {
	server *gin.Engine
	db     *mongo.Database
}

func NewRouter(server *gin.Engine, db *mongo.Database) *SetupRoutes {
	return &SetupRoutes{
		server: server,
		db:     db,
	}
}

func (r *SetupRoutes) setupRouters() {
	routers.SetupSessionRoutes(r.server, r.db)
}

func (r *SetupRoutes) Routers() {
	r.setupRouters()
}

func SetupAllRoutes(server *gin.Engine, db *mongo.Database) {
	router := NewRouter(server, db)
	router.Routers()
}
