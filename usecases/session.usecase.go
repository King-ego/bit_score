package usecases

import (
	"bit_score/entities"
	"bit_score/repositories"
)

type SessionUseCase interface {
	Login(username string, password string) (entities.Users, error)
}

type sessionUseCase struct {
	repository repositories.UsersRepository
}

func NewSessionUseCase(repository repositories.UsersRepository) SessionUseCase {
	return &sessionUseCase{
		repository: repository,
	}
}

func (s *sessionUseCase) Login(username string, password string) (entities.Users, error) {
	user, err := s.repository.GetByUserName(username)

	if err != nil {
		return entities.Users{}, err
	}

	return user, nil
}
