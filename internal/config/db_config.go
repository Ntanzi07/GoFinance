package config

import (
	"log"
	"os"

	"github.com/Ntanzi07/gofinance/internal/models"
	"github.com/joho/godotenv"
)

type DBConfig models.DBConfig

func LoadDBConfig() DBConfig {

	// Carrega o arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using system environment variables")
	}

	return DBConfig{
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASS"),
		DBHost: os.Getenv("DB_HOST"),
		Port:   os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_NAME"),
	}
}
