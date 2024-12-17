package infra

import (
	"log"

	"github.com/joho/godotenv"
)

func Initialize() {
	error := godotenv.Load()
	if error != nil {
		log.Fatal("Error loading .env file")
	}
}