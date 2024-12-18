package models

import (
	"time"

	"gorm.io/gorm"
)

type BasicModel struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Todo struct {
    BasicModel
    Title       string `gorm:"not null" json:"title"`
    IsCompleted bool   `gorm:"not null;default:false" json:"is_completed"`
    UserID      uint   `gorm:"not null" json:"user_id"`
}