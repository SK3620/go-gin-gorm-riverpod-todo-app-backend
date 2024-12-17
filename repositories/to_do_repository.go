package repositories

import "go-gin-gorm-riverpod-todo-app/models"


type ToDoInterface interface {
	FindAll() (*[]models.Todo, error)

}

// メモリ上のデータソースとして定義
type ToDoMemoryRepository struct {
	todos []models.Todo
}

func NewTodoMemoryRepository(todos []models.Todo) ToDoInterface {
	return &ToDoMemoryRepository{todos: todos}
}

func (r *ToDoMemoryRepository) FindAll() (*[]models.Todo, error) {
	return &r.todos, nil
}