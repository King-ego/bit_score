package usecases

import (
	"bit_score/entity"
	"bit_score/repositories"
)

type SessionUseCase interface {
	Login(username string, password string) (entity.Users, error)
}

type sessionUseCase struct {
	repository repositories.UsersRepository
}

func NewSessionUseCase(repository repositories.UsersRepository) SessionUseCase {
	return &sessionUseCase{
		repository: repository,
	}
}

func (s *sessionUseCase) Login(username string, password string) (entity.Users, error) {
	user, err := s.repository.GetByUserName(username)

	if err != nil {
		return entity.Users{}, err
	}

	return user, nil
}
