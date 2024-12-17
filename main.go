package main

import (
	"go-gin-gorm-riverpod-todo-app/controllers"
	"go-gin-gorm-riverpod-todo-app/infra"
	"go-gin-gorm-riverpod-todo-app/models"
	"go-gin-gorm-riverpod-todo-app/repositories"
	"go-gin-gorm-riverpod-todo-app/services"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()
	
	// サンプルデータを作成
	todos := []models.Todo{
		{ID: 1, Title: "タイトル１", IsCompleted: false},
		{ID: 2, Title: "タイトル２", IsCompleted: true},
		{ID: 3, Title: "タイトル３", IsCompleted: false},
	}

	todoRepository := repositories.NewTodoMemoryRepository(todos)
	todoService := services.NewTodoService(todoRepository)
	todoController := controllers.NewTodoController(todoService)

	r := gin.Default()
	r.GET("/todos", todoController.FindAll)
	r.POST("/todos", todoController.Create)
	r.PUT("/todos/:id", todoController.Update)
	r.DELETE("/todos/:id", todoController.Delete)
	r.Run()
}
