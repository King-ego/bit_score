package controllers

import (
	"bit_score/dto"
	"bit_score/usecases"

	"github.com/gin-gonic/gin"
)

type SessionsController struct {
	useCase usecases.SessionUseCase
}

func NewSessionsController(useCase usecases.SessionUseCase) *SessionsController {
	return &SessionsController{
		useCase: useCase,
	}
}

func (sc *SessionsController) CreateSession(c *gin.Context) {
	session := dto.SessionsDto{}

	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := sc.useCase.Login(session.UserName, session.Password)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to login"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Login successful",
		"user":    user,
		"token":   "dummy_token",
	})
}
