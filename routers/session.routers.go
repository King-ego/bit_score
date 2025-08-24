package routers

import (
	"bit_score/controllers"
	"bit_score/repositories"
	"bit_score/usecases"

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
	sessionRepository := repositories.NewUsersRepository(s.db)
	sessionUseCase := usecases.NewSessionUseCase(sessionRepository)
	sessionController := controllers.NewSessionsController(sessionUseCase)

	sessions := s.server.Group("/sessions")
	{
		sessions.POST("/", sessionController.CreateSession)
	}
}

func (s *SetupSessionRouters) Routers() {
	s.setupRouters()
}

func SetupSessionRoutes(server *gin.Engine, db *mongo.Database) {
	sessionRouter := NewSessionRouters(server, db)
	sessionRouter.Routers()
}
