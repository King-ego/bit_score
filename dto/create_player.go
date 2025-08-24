package dto

type CreatePlayerDto struct {
	Name string `json:"name" binding:"required"`
}
