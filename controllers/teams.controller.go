package controllers

import (
	"bit_score/dto"
	"bit_score/usecases"

	"github.com/gin-gonic/gin"
)

type TeamController struct {
	teamUseCase usecases.TeamUseCase
}

func NewTeamController(useCase usecases.TeamUseCase) *TeamController {
	return &TeamController{
		teamUseCase: useCase,
	}
}

func (tc *TeamController) CreateTeam(c *gin.Context) {
	team := dto.CreateTeamDto{}

	if err := c.ShouldBindJSON(&team); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := tc.teamUseCase.CreateTeam(team.Name, team.PrimaryColor, team.SecondColor)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create team"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Team created successfully",
		"team":    team,
	})
}

func (tc *TeamController) GetAllTeams(c *gin.Context) {
	teams, err := tc.teamUseCase.GetAllTeams()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get teams"})
		return
	}

	c.JSON(200, gin.H{
		"teams": teams,
	})
}

func (tc *TeamController) GetTeamByID(c *gin.Context) {
	id := c.Param("id")

	team, err := tc.teamUseCase.GetTeamByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Team not found!"})
		return
	}

	c.JSON(200, gin.H{
		"team": team,
	})
}

func (tc *TeamController) UpdateTeam(c *gin.Context) {
	id := c.Param("id")
	team := dto.UpdateTeamDto{}

	if err := c.ShouldBindJSON(&team); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := tc.teamUseCase.UpdateTeam(id, team.Name, team.PrimaryColor, team.SecondColor)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update team"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Team updated successfully",
	})
}

func (tc *TeamController) DeleteTeam(c *gin.Context) {
	id := c.Param("id")

	err := tc.teamUseCase.DeleteTeam(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete team"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Team deleted successfully",
	})
}
