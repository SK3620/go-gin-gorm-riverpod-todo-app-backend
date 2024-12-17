package controllers

import (
	"go-gin-gorm-riverpod-todo-app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoControllerInterface interface {
	FindAll(ctx *gin.Context)
}

type TodoController struct {
	service services.TodoServiceInterface
}

func NewTodoController(service services.TodoServiceInterface) TodoControllerInterface {
	return &TodoController{service: service}
}

func (c *TodoController) FindAll(ctx *gin.Context) {
	todos, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": todos})
}