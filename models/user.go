package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Username string `gorm:"not null" json:"username"`
    Email    string `gorm:"unique;not null" json:"email"`
    Password string `gorm:"not null" json:"password"`
    Todos    []Todo `gorm:"constraint:OnDelete:CASCADE" json:"todos"`
}
