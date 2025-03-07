package controllers

import (
	"go-gin-gorm-riverpod-todo-app/dto"
	"go-gin-gorm-riverpod-todo-app/models"
	"go-gin-gorm-riverpod-todo-app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ITodoController interface {
	FindAll(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type TodoController struct {
	service services.ITodoService
}

func NewTodoController(service services.ITodoService) ITodoController {
	// ITodoServiceは代入される具体的な値の型情報とポインタ情報を持つ

	// TodoController構造体はITodoControllerインターフェースを満たしており、ITodoServiceは具体的な値の型情報とポインタ情報を持つ
	return &TodoController{service: service}
}

// ===== TodoController構造体がITodoControllerインターフェースを満たすようにインターフェースに定義されたメソッドを実装する =====

// 全件取得
func (c *TodoController) FindAll(ctx *gin.Context) {
	user, exists := ctx.Get("user") // AuthMiddlewareにてSet()で設定した"user"キーで値を取り出す
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// userはany型のため、型アサーションを行う
	userId := user.(*models.User).ID

	todos, err := c.service.FindAll(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	// ネスト化したjsonレスポンスデータを送る
	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"todos": todos,
		},
	})
}

// 新規作成
func (c *TodoController) Create(ctx *gin.Context) {
	user, exists := ctx.Get("user") // AuthMiddlewareにてSet()で設定した"user"キーで値を取り出す
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// userはany型のため、型アサーションを行う
	userId := user.(*models.User).ID

	// リクエストデータを受け取る変数を用意
	var input dto.CreateToDoInput
	// リクエストデータをinput変数にバインド ここでリクエストデータのバリデーションが行われる
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTodo, err := c.service.Create(input, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": newTodo})
}

func (c *TodoController) Update(ctx *gin.Context) {
	user, exists := ctx.Get("user") // AuthMiddlewareにてSet()で設定した"user"キーで値を取り出す
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// userはany型のため、型アサーションを行う
	userId := user.(*models.User).ID

	todoId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	// リクエストデータを受け取る変数を用意
	var input dto.UpdateTodoInput
	// リクエストデータをinput変数にバインド ここでリクエストデータのバリデーションが行われる
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTodo, err := c.service.Update(uint(todoId), userId, input)
	if err != nil {
		if err.Error() == "Todo not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": updatedTodo})
}

func (c *TodoController) Delete(ctx *gin.Context) {
	user, exists := ctx.Get("user") // AuthMiddlewareにてSet()で設定した"user"キーで値を取り出す
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// userはany型のため、型アサーションを行う
	userId := user.(*models.User).ID

	todoId, error := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	error = c.service.Delete(uint(todoId), userId)
	if error != nil {
		if error.Error() == "Item not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": error.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": nil})
}