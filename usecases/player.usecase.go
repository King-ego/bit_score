package usecases

import "bit_score/repositories"

type PlayerUseCase struct {
	repository repositories.PlayersRepository
}

func NewPlayerRequest(repository repositories.PlayersRepository) PlayerUseCase {
	return &PlayerUseCase{repository: repository}
}

func (p *PlayerUseCase) CreatePlayer(name string) error {
	player := repositories.ICreatePlayer{
		Name: name,
	}
	err := p.repository.Create(player)
	if err != nil {
		return err
	}
	return nil
}
