package models

// DBConfig holds the database configuration details.
type DBConfig struct {
	DBUser string
	DBPass string
	DBHost string
	DBName string
	Port   string
}

// Notes:
// - Port is stored as string to match environment variables; convert to int
//   if you need numeric operations.
// - DBConfig is mapped from environment variables in `internal/config`.
