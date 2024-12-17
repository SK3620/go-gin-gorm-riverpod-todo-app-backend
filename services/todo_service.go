package services

import (
	"go-gin-gorm-riverpod-todo-app/models"
	"go-gin-gorm-riverpod-todo-app/repositories"
)

type TodoServiceInterface interface {
	FindAll() (*[]models.Todo, error)
}

type TodoService struct {
	repository repositories.ToDoInterface
}

func NewTodoService(repository repositories.ToDoInterface) TodoServiceInterface {
	return &TodoService{repository: repository}
}

func (s *TodoService) FindAll() (*[]models.Todo, error) {
	return s.repository.FindAll()
}