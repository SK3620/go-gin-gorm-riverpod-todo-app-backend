package repositories

import (
	"errors"
	"go-gin-gorm-riverpod-todo-app/models"

	"gorm.io/gorm"
)

type ITodoRepository interface {
	FindAll(userId uint) (*[]models.Todo, error)
	FindById(todoId uint, userId uint) (*models.Todo, error)
	Create(newTodo models.Todo) (*models.Todo, error)
	Update(updateTodo models.Todo) (*models.Todo, error)
	Delete(todoId uint, userId uint) error
}

type TodoMemoryRepository struct {
	todos []models.Todo
}

func NewTodoMemoryRepository(todos []models.Todo) ITodoRepository {
	return &TodoMemoryRepository{todos: todos}
}

func (r *TodoMemoryRepository) FindAll(userId uint) (*[]models.Todo, error) {
	return &r.todos, nil
}

func (r *TodoMemoryRepository) FindById(todoId uint, userId uint) (*models.Todo, error) {
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

func (r *TodoMemoryRepository) Delete(todoId uint, userId uint) error {
	for i, v := range r.todos {
		if v.ID == todoId {
			r.todos = append(r.todos[:i], r.todos[i+1:]...)
			return nil
		}
	}
	return errors.New("Todo not found")
}

// ================================================================================================================================================================

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) ITodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) Create(newTodo models.Todo) (*models.Todo, error) {
	result := r.db.Create(&newTodo)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newTodo, nil
}

func (r *TodoRepository) Delete(todoId uint, userId uint) error {
	deleteTodo, error := r.FindById(todoId, userId)
	if error != nil {
		return error
	}

	result := r.db.Delete(&deleteTodo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *TodoRepository) FindAll(userId uint) (*[]models.Todo, error) {
	var todos []models.Todo
	result := r.db.Find(&todos, "user_id = ?", userId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &todos, nil
}

func (r *TodoRepository) FindById(todoId uint, userId uint) (*models.Todo, error) {
	var todo models.Todo
	result := r.db.First(&todo, "id = ? AND user_id = ?", todoId, userId)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("Todo not found")
		}
		return nil, result.Error
	}
	return &todo, nil
}

func (r *TodoRepository) Update(updateTodo models.Todo) (*models.Todo, error) {
	result := r.db.Save(&updateTodo)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updateTodo, nil
}