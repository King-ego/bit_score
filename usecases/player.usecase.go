package usecases

import "bit_score/repositories"

type LoginRequest struct {
	repository repositories.PlayersRepository
}
