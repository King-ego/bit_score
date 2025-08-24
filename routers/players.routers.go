package routers

import (
	"bit_score/controllers"
	"bit_score/repositories"
	"bit_score/usecases"

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
	playerRepository := repositories.NewPlayersRepository(p.db)
	playerUseCase := usecases.NewPlayerRequest(playerRepository)
	playerController := controllers.NewPlayerController(playerUseCase)

	player := p.server.Group("/players")
	{
		player.POST("/", playerController.CreatePlayer)
	}

}

func (p *SetupPlayersRouter) Routers() {
	p.setupRouters()
}

func SetupPlayersRoutes(server *gin.Engine, db *mongo.Database) {
	playersRouter := NewPlayersRouters(server, db)
	playersRouter.Routers()
}
