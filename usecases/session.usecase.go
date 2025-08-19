package usecases

type SessionUseCase interface {
	Login(username string, password string) (string, error)
}

type sessionUseCase struct {
}

func NewSessionUseCase() SessionUseCase {
	return &sessionUseCase{}
}

func (s *sessionUseCase) Login(username string, password string) (string, error) {
	return "Success", nil
}
