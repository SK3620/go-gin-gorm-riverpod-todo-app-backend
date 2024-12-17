package main

import (
	"go-gin-gorm-riverpod-todo-app/controllers"
	"go-gin-gorm-riverpod-todo-app/infra"

	// "go-gin-gorm-riverpod-todo-app/models"
	"go-gin-gorm-riverpod-todo-app/repositories"
	"go-gin-gorm-riverpod-todo-app/services"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()
	
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

	r := gin.Default()
	todoRouter := r.Group("/todos")
	todoRouter.GET("", todoController.FindAll)
	todoRouter.POST("", todoController.Create)
	todoRouter.PUT("/:id", todoController.Update)
	todoRouter.DELETE("/:id", todoController.Delete)
	r.Run("localhost:8080")
}
