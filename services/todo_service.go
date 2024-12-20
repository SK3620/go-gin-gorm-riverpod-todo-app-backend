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

// ファクトリ関数
func NewTodoService(repository repositories.ITodoRepository) ITodoService {
	// ITodoRepositoryは代入される具体的な値の型情報とポインタ情報を持つ

	// TodoService構造体はITodoServiceインターフェースを満たしており、ITodoServiceは具体的な値の型情報とポインタ情報を持つ
	return &TodoService{repository: repository}
}

// ===== TodoService構造体がITodoServiceインターフェースを満たすようにインターフェースに定義されたメソッドを実装する =====

// 全件取得
func (s *TodoService) FindAll(userId uint) (*[]models.Todo, error) {
	return s.repository.FindAll(userId)
}

// idに一致するデータを取得
func (s *TodoService) FindById(todoId uint, userId uint) (*models.Todo, error) {
	return s.repository.FindById(todoId, userId)
}

// 新規作成
func (s *TodoService) Create(createTodoInput dto.CreateToDoInput, userId uint) (*models.Todo, error) {
	newTodo := models.Todo{
		Title: createTodoInput.Title,
		IsCompleted: false,
		UserID: userId,
	}
	return s.repository.Create(newTodo)
}

// 更新
func (s *TodoService) Update(todoId uint, userId uint, updateTodoInput dto.UpdateTodoInput) (*models.Todo, error) {
	targetItem, error := s.FindById(todoId, userId)
	if error != nil {
		return nil, error
	}

	// 対象のtodoのタイトルを更新
	if updateTodoInput.Title != nil {
		targetItem.Title = *updateTodoInput.Title
	}
	// 対象のtodoの完了フラグを更新
	if updateTodoInput.IsCompleted != nil {
		targetItem.IsCompleted = *updateTodoInput.IsCompleted
	}
	return s.repository.Update(*targetItem)
}

// 削除
func (s *TodoService) Delete(todoId uint, userId uint) error {
	return s.repository.Delete(todoId, userId)
}