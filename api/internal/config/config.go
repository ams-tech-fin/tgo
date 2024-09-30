package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {

	if err := godotenv.Load(); err != nil {
		log.Println("File .env not found")
	}
}

func GetEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
