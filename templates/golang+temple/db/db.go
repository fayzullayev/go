package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {

	connStr := "postgres://postgres:12345@localhost/projects_management?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
