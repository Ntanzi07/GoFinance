package database

import (
	"database/sql"
	"fmt"

	"github.com/Ntanzi07/gofinance/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

// Connect establishes a connection to the MariaDB database.
func Connect() (*sql.DB, error) {

	DBConfig := config.LoadDBConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		DBConfig.DBUser,
		DBConfig.DBPass,
		DBConfig.DBHost,
		DBConfig.Port,
		DBConfig.DBName,
	)

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
