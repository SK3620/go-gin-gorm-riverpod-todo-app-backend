package main

import (
	"go-gin-gorm-riverpod-todo-app/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// サンプルデータを作成
	todos := []models.Todo{
		{ID: 1, Title: "タイトル１", IsCompleted: false},
		{ID: 2, Title: "タイトル２", IsCompleted: true},
		{ID: 3, Title: "タイトル３", IsCompleted: false},
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 0.0.0.0:8080 でサーバーを立てます。
}
