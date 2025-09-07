package routers

import (
	"bit_score/controllers"
	"bit_score/repositories"
	"bit_score/usecases"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type SetupTeamsRouter struct {
	server *gin.Engine
	db     *mongo.Database
}

func NewTeamsRouters(server *gin.Engine, db *mongo.Database) *SetupTeamsRouter {
	return &SetupTeamsRouter{
		server: server,
		db:     db,
	}
}

func (t *SetupTeamsRouter) setupRouters() {
	teamRepository := repositories.NewTeamsRepository(t.db)
	teamUseCase := usecases.NewTeamRequest(teamRepository)
	teamController := controllers.NewTeamController(teamUseCase)

	teams := t.server.Group("/teams")
	{
		teams.POST("/", teamController.CreateTeam)
		teams.GET("/", teamController.GetAllTeams)
		teams.GET("/:id", teamController.GetTeamByID)
		teams.PUT("/:id", teamController.UpdateTeam)
		teams.DELETE("/:id", teamController.DeleteTeam)
	}
}

func (t *SetupTeamsRouter) Routers() {
	t.setupRouters()
}

func SetupTeamsRoutes(server *gin.Engine, db *mongo.Database) {
	teamsRouter := NewTeamsRouters(server, db)
	teamsRouter.Routers()
}
