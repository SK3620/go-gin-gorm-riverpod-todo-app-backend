package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title string `gorm:"not null"`
	IsCompleted bool `gorm:"not null;default:false"`
	UserID uint `gorm:"not null"`
}