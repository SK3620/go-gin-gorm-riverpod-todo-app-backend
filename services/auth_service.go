package services

import (
	"go-gin-gorm-riverpod-todo-app/models"
	"go-gin-gorm-riverpod-todo-app/repositories"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	SignUp(username string, email string, password string) error
	Login(email string, password string) (*string, error)
}

type AuthService struct {
	repository repositories.IAuthRepository
}

func NewAuthService(respository repositories.IAuthRepository) IAuthService {
	return &AuthService{repository: respository}
}

func (s *AuthService) SignUp(username string, email string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Username: username,
		Email: email,
		Password: string(hashedPassword),
	}
	return s.repository.CreateUser(user)
}

func (s *AuthService) Login(email string, password string) (*string, error) {
	foundUser, err := s.repository.FindUser(email)
	if err != nil {
		return nil,  err
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password))
	if  err != nil {
		return nil,  err
	}

	token, err := CreateToken(foundUser.ID, foundUser.Email)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func CreateToken(userId uint, email string) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"email": email,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	// 秘密鍵を使用して著名を行う
	tokenString, error := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if error != nil {
		return nil, error
	}
	return &tokenString, nil
}