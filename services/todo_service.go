package services

import (
	"go-gin-gorm-riverpod-todo-app/dto"
	"go-gin-gorm-riverpod-todo-app/models"
	"go-gin-gorm-riverpod-todo-app/repositories"
)

type ITodoService interface {
	FindAll(userId uint) (*[]models.Todo, error)
	FindById(todoId uint, userId uint) (*models.Todo, error)
	Create(createTodoInput dto.CreateToDoInput, userId uint) (*models.Todo, error)
	Update(todoId uint, userId uint, updateItemInput dto.UpdateTodoInput) (*models.Todo, error)
	Delete(todoId uint, userId uint) error
}

type TodoService struct {
	repository repositories.ITodoRepository
}

func NewTodoService(repository repositories.ITodoRepository) ITodoService {
	return &TodoService{repository: repository}
}

func (s *TodoService) FindAll(userId uint) (*[]models.Todo, error) {
	return s.repository.FindAll(userId)
}

func (s *TodoService) FindById(todoId uint, userId uint) (*models.Todo, error) {
	return s.repository.FindById(todoId, userId)
}

func (s *TodoService) Create(createTodoInput dto.CreateToDoInput, userId uint) (*models.Todo, error) {
	newTodo := models.Todo{
		Title: createTodoInput.Title,
		IsCompleted: false,
		UserID: userId,
	}
	return s.repository.Create(newTodo)
}

func (s *TodoService) Update(todoId uint, userId uint, updateTodoInput dto.UpdateTodoInput) (*models.Todo, error) {
	targetItem, error := s.FindById(todoId, userId)
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

func (s *TodoService) Delete(todoId uint, userId uint) error {
	return s.repository.Delete(todoId, userId)
}