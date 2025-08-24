package controllers

import (
	"bit_score/dto"
	"bit_score/usecases"

	"github.com/gin-gonic/gin"
)

type PlayerController struct {
	playerUseCase usecases.PlayerUseCase
}

func NewPlayerController(useCase usecases.PlayerUseCase) *PlayerController {
	return &PlayerController{
		playerUseCase: useCase,
	}
}

func (sc *PlayerController) CreatePlayer(c *gin.Context) {
	player := dto.CreatePlayerDto{}

	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := sc.playerUseCase.CreatePlayer(player.Name)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create player"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Player created successfully",
		"player":  player,
	})
}
