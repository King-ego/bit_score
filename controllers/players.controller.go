package controllers

import (
	"bit_score/dto"

	"github.com/gin-gonic/gin"
)

type PlayerController struct {
}

func NewSessionController() *PlayerController {
	return &PlayerController{}
}

func (sc *PlayerController) CreatePlayer(c *gin.Context) {
	player := dto.CreatePlayerDto{}

	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}
