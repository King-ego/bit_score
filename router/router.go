package router

import (
	"bit_score/routers"

	"github.com/gin-gonic/gin"
)

type SetupRoutes struct {
	server *gin.Engine
}

func NewRouter(server *gin.Engine) *SetupRoutes {
	return &SetupRoutes{
		server: server,
	}
}

func (r *SetupRoutes) setupRouters() {
	routers.SetupSessionRoutes(r.server)
}

func (r *SetupRoutes) Routers() {
	r.setupRouters()
}

func SetupAllRoutes(server *gin.Engine) {
	router := NewRouter(server)
	router.Routers()
}
