package main

import (
	"go-gin-gorm-riverpod-todo-app/infra"
	"go-gin-gorm-riverpod-todo-app/models"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	if error := db.AutoMigrate(&models.Todo{}, &models.User{}); error != nil {
		panic("Failed to migrate database")
	}
}