package services

import (
	"go-gin-gorm-riverpod-todo-app/models"
	"go-gin-gorm-riverpod-todo-app/repositories"

	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	SignUp(username string, email string, password string) error
}

type AuthService struct {
	repository repositories.IAuthRepository
}

func NewAuthService(respository repositories.IAuthRepository) IAuthService {
	return &AuthService{repository: respository}
}

func (s *AuthService) SignUp(username string, email string, password string) error {
	hashedPassword, error := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if error != nil {
		return error
	}

	user := models.User{
		Username: username,
		Email: email,
		Password: string(hashedPassword),
	}
	return s.repository.CreateUser(user)
}