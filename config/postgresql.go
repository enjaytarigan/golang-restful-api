package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	PG_USER := os.Getenv("PG_USER")
	PG_HOST := os.Getenv("PG_HOST")
	PG_PASSWORD := os.Getenv("PG_PASSWORD")
	PG_PORT := os.Getenv("PG_PORT")
	PG_DATABASE := os.Getenv("PG_DATABASE")
	PG_SSLMODE := os.Getenv("PG_SSLMODE")

	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", PG_USER, PG_PASSWORD, PG_HOST, PG_PORT, PG_DATABASE, PG_SSLMODE)

	db, err := sql.Open("postgres", connection)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
