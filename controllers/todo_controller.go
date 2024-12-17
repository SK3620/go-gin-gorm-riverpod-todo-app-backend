package controllers

import (
	"go-gin-gorm-riverpod-todo-app/dto"
	"go-gin-gorm-riverpod-todo-app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ITodoController interface {
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type TodoController struct {
	service services.ITodoService
}

func NewTodoController(service services.ITodoService) ITodoController {
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

func (c *TodoController) FindById(ctx *gin.Context) {
	todoId, error := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	todo, error := c.service.FindById(uint(todoId))
	if error != nil {
		if error.Error() == "Todo not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": error.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{"data": todo})
}

func (c *TodoController) Create(ctx *gin.Context) {
	var input dto.CreateToDoInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTodo, err := c.service.Create(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": newTodo})
}

func (c *TodoController) Update(ctx *gin.Context) {
	todoId, error := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	var input dto.UpdateTodoInput
	if error := ctx.ShouldBindJSON(&input); error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}

	updatedTodo, error := c.service.Update(uint(todoId), input)
	if error != nil {
		if error.Error() == "Todo not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": error.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": updatedTodo})

}

func (c *TodoController) Delete(ctx *gin.Context) {
	todoId, error := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	error = c.service.Delete(uint(todoId))
	if error != nil {
		if error.Error() == "Item not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": error.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.Status(http.StatusOK)
}