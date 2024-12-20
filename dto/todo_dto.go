package dto

type CreateToDoInput struct {
	Title string `json:"title" binding:"required"`
}

// ポインタ型（*）でnullを許容
// omitnilでnilの場合はそのフィールドのバリデーションはしない
type UpdateTodoInput struct {
	Title *string `json:"title" binding:"omitnil"`
	IsCompleted *bool `json:"is_completed" binding:"omitnil"`
}