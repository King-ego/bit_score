package routers

import (
	"bit_score/controllers"

	"github.com/gin-gonic/gin"
)

type SetupSessionRouters struct {
	server *gin.Engine
}

func NewSessionRouters(server *gin.Engine) *SetupSessionRouters {
	return &SetupSessionRouters{
		server: server,
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

func SetupSessionRoutes(server *gin.Engine) {
	sessionRouter := NewSessionRouters(server)
	sessionRouter.Routers()
}
