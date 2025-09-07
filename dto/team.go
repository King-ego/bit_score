package dto

type CreateTeamDto struct {
	Name         string `json:"name" binding:"required"`
	PrimaryColor string `json:"primary_color" binding:"required"`
	SecondColor  string `json:"second_color" binding:"required"`
}

type UpdateTeamDto struct {
	Name         string `json:"name"`
	PrimaryColor string `json:"primary_color"`
	SecondColor  string `json:"second_color"`
}
