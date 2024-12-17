package repositories

import (
	"errors"
	"go-gin-gorm-riverpod-todo-app/models"
)


type ITodoRepository interface {
	FindAll() (*[]models.Todo, error)
	FindById(todoId uint) (*models.Todo, error)
	Create(newTodo models.Todo) (*models.Todo, error)
	Update(updateTodo models.Todo) (*models.Todo, error)
	Delete(todoId uint) error
}

type TodoMemoryRepository struct {
	todos []models.Todo
}

func NewTodoMemoryRepository(todos []models.Todo) ITodoRepository {
	return &TodoMemoryRepository{todos: todos}
}

func (r *TodoMemoryRepository) FindAll() (*[]models.Todo, error) {
	return &r.todos, nil
}

func (r *TodoMemoryRepository) FindById(todoId uint) (*models.Todo, error) {
	for _, v := range r.todos {
		if v.ID == todoId {
			return &v, nil
		}
	}
	return nil, errors.New("Todo not found")
}

func (r *TodoMemoryRepository) Create(newItem models.Todo) (*models.Todo, error) {
	newItem.ID = uint(len(r.todos) + 1)
	r.todos = append(r.todos, newItem)
	return &newItem, nil
}

func (r *TodoMemoryRepository) Update(updateTodo models.Todo) (*models.Todo, error) {
	for i, v := range r.todos {
		if v.ID == updateTodo.ID {
			r.todos[i] = updateTodo
			return &r.todos[i], nil
		}
	}
	return nil, errors.New("Unexpected error")
}

func (r *TodoMemoryRepository) Delete(todoId uint) error {
	for i, v := range r.todos {
		if v.ID == todoId {
			r.todos = append(r.todos[:i], r.todos[i+1:]...)
			return nil
		}
	}
	return errors.New("Item not found")
}