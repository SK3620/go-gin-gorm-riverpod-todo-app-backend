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

// ↑ メモリ上でデータを操作
// =================================================================================================================================================================================
// ↓ DBを用いてデータ操作

type TodoRepository struct {
	db *gorm.DB
}

// ファクトリ関数
func NewTodoRepository(db *gorm.DB) ITodoRepository {
	// ToDoRepository構造体がITodoRepositoryインターフェースを満たしており、ITodoRepositoryは具体的な値の型情報とポインタ情報を持つ
	return &TodoRepository{db: db}
}

// ===== ToDoRepository構造体がITodoRepositoryインターフェースを満たすようにインターフェースに定義されたメソッドを実装する =====

// 新規作成
func (r *TodoRepository) Create(newTodo models.Todo) (*models.Todo, error) {
	result := r.db.Create(&newTodo) // 参照を渡す
	if result.Error != nil {
		return nil, result.Error
	}
	return &newTodo, nil
}

// 削除
func (r *TodoRepository) Delete(todoId uint, userId uint) error {
	deleteTodo, error := r.FindById(todoId, userId)
	if error != nil {
		return error
	}

	result := r.db.Delete(&deleteTodo)
	if result.Error != nil {
		return result.Error
	}
	return nil // 削除成功でnilを返す
}

// 全件取得
func (r *TodoRepository) FindAll(userId uint) (*[]models.Todo, error) {
	var todos []models.Todo
	result := r.db.Find(&todos, "user_id = ?", userId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &todos, nil
}

// idに一致するデータを取得
func (r *TodoRepository) FindById(todoId uint, userId uint) (*models.Todo, error) {
	// models.Todo型の値を格納する変数を定義
	var todo models.Todo

	// Fisrt()で最初にヒットした一件のみ取得
	// 第一引数には検索結果を格納するモデルの参照、第二引数には検索条件を渡す
	result := r.db.First(&todo, "id = ? AND user_id = ?", todoId, userId)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("Todo not found")
		}
		return nil, result.Error
	}
	return &todo, nil
}

// 更新
func (r *TodoRepository) Update(updateTodo models.Todo) (*models.Todo, error) {
	result := r.db.Save(&updateTodo)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updateTodo, nil
}