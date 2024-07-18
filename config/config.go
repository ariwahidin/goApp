package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DBUser     string
	DBPassword string
	DBName     string
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
}
