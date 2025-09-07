package usecases

import (
	"bit_score/entities"
	"bit_score/repositories"
)

type TeamUseCase interface {
	CreateTeam(name, primaryColor, secondColor string) error
	GetAllTeams() ([]entities.Teams, error)
	GetTeamByID(id string) (*entities.Teams, error)
	UpdateTeam(id, name, primaryColor, secondColor string) error
	DeleteTeam(id string) error
}

type teamUseCase struct {
	repository repositories.TeamsRepository
}

func NewTeamRequest(repository repositories.TeamsRepository) TeamUseCase {
	return &teamUseCase{repository: repository}
}

func (t *teamUseCase) CreateTeam(name, primaryColor, secondColor string) error {
	team := repositories.ICreateTeam{
		Name:         name,
		PrimaryColor: primaryColor,
		SecondColor:  secondColor,
	}
	err := t.repository.Create(team)
	if err != nil {
		return err
	}
	return nil
}

func (t *teamUseCase) GetAllTeams() ([]entities.Teams, error) {
	teams, err := t.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return teams, nil
}

func (t *teamUseCase) GetTeamByID(id string) (*entities.Teams, error) {
	team, err := t.repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return team, nil
}

func (t *teamUseCase) UpdateTeam(id, name, primaryColor, secondColor string) error {
	team := repositories.IUpdateTeam{
		Name:         name,
		PrimaryColor: primaryColor,
		SecondColor:  secondColor,
	}
	err := t.repository.Update(id, team)
	if err != nil {
		return err
	}
	return nil
}

func (t *teamUseCase) DeleteTeam(id string) error {
	err := t.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
