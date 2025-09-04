package usecases

import "bit_score/repositories"

type PlayerUseCase interface {
	CreatePlayer(name string) error
}

type playerUseCase struct {
	repository repositories.PlayersRepository
}

func NewPlayerRequest(repository repositories.PlayersRepository) PlayerUseCase {
	return &playerUseCase{repository: repository}
}

func (p *playerUseCase) CreatePlayer(name string) error {
	player := repositories.ICreatePlayer{
		Name: name,
	}
	err := p.repository.Create(player)
	if err != nil {
		return err
	}
	return nil
}
