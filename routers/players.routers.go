package routers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type SetupPlayersRouter struct {
	server *gin.Engine
	db     *mongo.Database
}

func NewPlayersRouters(server *gin.Engine, db *mongo.Database) *SetupPlayersRouter {
	return &SetupPlayersRouter{
		server: server,
		db:     db,
	}
}

func (p *SetupPlayersRouter) setupRouters() {

}

func (p *SetupPlayersRouter) Routers() {
	p.setupRouters()
}

func SetupPlayersRoutes(server *gin.Engine, db *mongo.Database) {
	playersRouter := NewPlayersRouters(server, db)
	playersRouter.Routers()
}
