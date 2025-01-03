package services

import (
	"fmt"
	"go-gin-gorm-riverpod-todo-app/models"
	"go-gin-gorm-riverpod-todo-app/repositories"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	SignUp(username string, email string, password string) (*string, error)
	Login(email string, password string) (*string, error)
	GetUserFromToken(tokenString string) (*models.User, error)
}

type AuthService struct {
	repository repositories.IAuthRepository
}

func NewAuthService(respository repositories.IAuthRepository) IAuthService {
	return &AuthService{repository: respository}
}

func (s *AuthService) SignUp(username string, email string, password string) (*string, error) {
	// パスワードのハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Username: username,
		Email: email,
		Password: string(hashedPassword),
	}
	createUserErr := s.repository.CreateUser(user)
	if createUserErr != nil {
		return nil, createUserErr
	}

	foundUser, err := s.repository.FindUser(email)
	if err != nil {
		return nil,  err
	}

	// JWTTokenを生成
	token, err := CreateToken(foundUser.ID, foundUser.Email)
	if err != nil {
		return nil, err
	}

	return token,  nil
}

func (s *AuthService) Login(email string, password string) (*string, error) {
	foundUser, err := s.repository.FindUser(email)
	if err != nil {
		return nil,  err
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password))
	// パスワードが一致しない場合
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
	// tokenの生成
	// Claimsはtokenに含める様々な情報を指す
	// sub（subject）→ ユーザー識別子
	// exp → tokenの有効期限
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"email": email,
		"exp": time.Now().Add(time.Hour).Unix(), // 生成から1時間後に期限を設定
	})

	// 秘密鍵を使用して著名を行う
	tokenString, error := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if error != nil {
		return nil, error
	}
	return &tokenString, nil
}

// トークンに含まれる情報を基にユーザー情報を取得
func (s *AuthService) GetUserFromToken(tokenString string) (*models.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// トークンの署名を検証するコールバック関数内で、トークンの解析を行う

		// 署名がHMACか否か確認
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	var user *models.User
	// token.Claims: トークン内に含まれるデータ
	// クレームが正しい形式（MapClaims）であるか確認
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return nil, jwt.ErrTokenExpired
		}
		// メールアドレスを基にデータベースからユーザー情報を取得
		user, err = s.repository.FindUser(claims["email"].(string))
		if err != nil {
			return nil, err
		}
	}
	return user, nil
}