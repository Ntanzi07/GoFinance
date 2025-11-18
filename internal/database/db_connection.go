package database

import (
	"database/sql"
	"fmt"

	"github.com/Ntanzi07/gofinance/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

// Connect establishes a connection to the MariaDB database.
//
// It constructs the DSN from environment configuration (via config.LoadDBConfig()),
// opens a connection pool and verifies connectivity with Ping. The returned
// *sql.DB should be reused across the application and closed on shutdown.
func Connect() (*sql.DB, error) {

	DBConfig := config.LoadDBConfig()

	// Build DSN: user:pass@tcp(host:port)/dbname
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		DBConfig.DBUser,
		DBConfig.DBPass,
		DBConfig.DBHost,
		DBConfig.Port,
		DBConfig.DBName,
	)

	// sql.Open does not establish a connection immediately; it prepares the
	// database handle (connection pool). Ping verifies the DB is reachable.
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
