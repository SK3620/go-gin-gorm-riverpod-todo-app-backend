package main

import (
	"go-gin-gorm-riverpod-todo-app/controllers"
	"go-gin-gorm-riverpod-todo-app/infra"
	"go-gin-gorm-riverpod-todo-app/middlwares"

	// "go-gin-gorm-riverpod-todo-app/models"
	"go-gin-gorm-riverpod-todo-app/repositories"
	"go-gin-gorm-riverpod-todo-app/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	// // サンプルデータを作成
	// todos := []models.Todo{
	// 	{ID: 1, Title: "タイトル１", IsCompleted: false},
	// 	{ID: 2, Title: "タイトル２", IsCompleted: true},
	// 	{ID: 3, Title: "タイトル３", IsCompleted: false},
	// }

	// todoRepository := repositories.NewTodoMemoryRepository(todos)
	todoRepository := repositories.NewTodoRepository(db)
	todoService := services.NewTodoService(todoRepository)
	todoController := controllers.NewTodoController(todoService)

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	r := gin.Default()
	authRouter := r.Group("/auth")
	todoRouterWithAuth := r.Group("/todos", middlwares.AuthMiddlware(authService))

	authRouter.POST("/sign_up", authController.SignUp)
	authRouter.POST("/login", authController.Login)

	todoRouterWithAuth.GET("", todoController.FindAll)
	todoRouterWithAuth.POST("", todoController.Create)
	todoRouterWithAuth.PUT("/:id", todoController.Update)
	todoRouterWithAuth.DELETE("/:id", todoController.Delete)

	return r
}

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	r := setupRouter(db)
	
	r.Run("localhost:8080")
}
