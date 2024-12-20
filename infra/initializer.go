package infra

import (
	"log"

	"github.com/joho/godotenv"
)

func Initialize() {
	error := godotenv.Load() // 引数にファイル名を渡す（省略するとデフォルトで.envファイルを読み込み）
	if error != nil {
		log.Fatal("Error loading .env file")
	}
}