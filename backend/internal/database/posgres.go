package database

import (
	"database/sql"
	"os"
)

func NewPostgres() (*sql.DB, error) {
	dsn := os.Getenv("DATABASE_URL")

	return sql.Open("postgres", dsn)
}
