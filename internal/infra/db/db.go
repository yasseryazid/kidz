package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// OpenFromEnv opens a postgres connection using DB_* env vars.
func OpenFromEnv() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	if host == "" || port == "" || name == "" || user == "" {
		return nil, fmt.Errorf("missing DB_* env vars")
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, name)
	return sql.Open("pgx", dsn)
}
