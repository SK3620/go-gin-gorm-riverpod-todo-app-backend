package controllers

import (
	"go-gin-gorm-riverpod-todo-app/dto"
	"go-gin-gorm-riverpod-todo-app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IAuthController interface{
	SignUp(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type AuthController struct {
	service services.IAuthService
}

func NewAuthController(service services.IAuthService) IAuthController {
	return &AuthController{service: service}
}

// サインアップ
func (c *AuthController) SignUp(ctx *gin.Context) {
	// リクエストデータを格納する変数を用意
	var input dto.SignUpInput

	 // リクエストデータをinput変数にバインド ここでリクエストデータのバリデーションが行わる
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := c.service.SignUp(input.Usermame, input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// ネスト化したjsonレスポンスデータを返す
	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"username": input.Usermame,
			"email":    input.Email,
			"password": input.Password,
			"jwt_token": token, // 生成したJWTToken
		},
	})
}

// ログイン
func (c *AuthController) Login(ctx *gin.Context) {
	// リクエストデータを格納する変数を用意
	var input dto.LoginInput

	// リクエストデータをinput変数にバインド ここでリクエストデータのバリデーションが行わる
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := c.service.Login(input.Email, input.Password)
	if err != nil {
		if err.Error() == "User not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ネスト化したjsonレスポンスデータを返す
	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"username": "",
			"email":    input.Email,
			"password": input.Password,
			"jwt_token": token, // 生成したJWTToken
		},
	})
}