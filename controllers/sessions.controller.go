package controllers

import (
	"bit_score/dto"

	"github.com/gin-gonic/gin"
)

type SessionsController struct {
}

func NewSessionsController() *SessionsController {
	return &SessionsController{}
}

func (sc *SessionsController) CreateSession(c *gin.Context) {
	session := dto.SessionsDto{}

	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}
