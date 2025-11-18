package config

import (
	"log"
	"os"

	"github.com/Ntanzi07/gofinance/internal/models"
	"github.com/joho/godotenv"
)

// DBConfig is an alias for models.DBConfig used by the config package.
type DBConfig models.DBConfig

// LoadDBConfig loads database configuration from environment variables.
// It tries to load a local .env file first (useful in development).
func LoadDBConfig() DBConfig {

	// Attempt to load .env (no error if file missing; fall back to system env)
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using system environment variables")
	}

	// Map environment variables into DBConfig
	return DBConfig{
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASS"),
		DBHost: os.Getenv("DB_HOST"),
		Port:   os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_NAME"),
	}
}

// LoadJwt returns the JWT secret as a byte slice (used by JWT middleware).
func LoadJwt() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}
