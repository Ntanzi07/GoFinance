package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Connect() (*sql.DB, error) {
	// Carrega o arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using system environment variables")
	}

	// Pega as variáveis
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("✅ Connected to MariaDB!")
	return db, nil
}
