package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Env(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)
}
