package infra

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	// dsnを生成するため、環境変数から取得
	dsn := fmt.Sprintf(
		// %s はプレースホルダーで、順番に引数から値が挿入される
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	
	// DBへの接続
	db, error := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if error != nil {
		panic("Failed to connect database")
	}

	return db
}