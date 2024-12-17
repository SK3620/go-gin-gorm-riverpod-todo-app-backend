package services

import (
	"go-gin-gorm-riverpod-todo-app/dto"
	"go-gin-gorm-riverpod-todo-app/models"
	"go-gin-gorm-riverpod-todo-app/repositories"
)

type ITodoService interface {
	FindAll() (*[]models.Todo, error)
	FindById(todoId uint) (*models.Todo, error)
	Create(createTodoInput dto.CreateToDoInput) (*models.Todo, error)
	Update(todoId uint, updateItemInput dto.UpdateTodoInput) (*models.Todo, error)
	Delete(todoId uint) error
}

type TodoService struct {
	repository repositories.ITodoRepository
}

func NewTodoService(repository repositories.ITodoRepository) ITodoService {
	return &TodoService{repository: repository}
}

func (s *TodoService) FindAll() (*[]models.Todo, error) {
	return s.repository.FindAll()
}

func (s *TodoService) FindById(todoId uint) (*models.Todo, error) {
	return s.repository.FindById(todoId)
}

func (s *TodoService) Create(createTodoInput dto.CreateToDoInput) (*models.Todo, error) {
	newTodo := models.Todo{
		Title: createTodoInput.Title,
		IsCompleted: false,
	}
	return s.repository.Create(newTodo)
}

func (s *TodoService) Update(todoId uint, updateTodoInput dto.UpdateTodoInput) (*models.Todo, error) {
	targetItem, error := s.FindById(todoId)
	if error != nil {
		return nil, error
	}

	if updateTodoInput.Title != nil {
		targetItem.Title = *updateTodoInput.Title
	}
	if updateTodoInput.IsCompleted != nil {
		targetItem.IsCompleted = *updateTodoInput.IsCompleted
	}
	return s.repository.Update(*targetItem)
}

func (s *TodoService) Delete(todoId uint) error {
	return s.repository.Delete(todoId)
}