package auth

import (
	"clinica-api/internal/models"
	"bytes"
	"fmt"
)

type AuthService interface {
	Login(username string, password [32]byte) (models.Login, error)
	GetUsuario(username string) (models.User, error)
}

type authService struct {
	repo AuthRepository
}

func NewAuthService(repo AuthRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Login(username string, password [32]byte) (models.Login, error) {
	login, err := s.repo.Autenticar(username)
	if err != nil {
		return models.Login{}, err
	}
	if !bytes.Equal(login.PasswordCriptografado, password[:]) {
		return login, fmt.Errorf("senha incorreta")
	}

	return login, nil
}

func (s *authService) GetUsuario(username string) (models.User, error) {
	user, err := s.repo.GetUsuario(username)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
