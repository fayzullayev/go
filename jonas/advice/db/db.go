package db

import (
	"database/sql"
	"fmt"
)
import _ "github.com/lib/pq"

var DB *sql.DB

func InitDB() *sql.DB {
	var err error

	DB, err = sql.Open("postgres", "postgres://postgres:12345@localhost:5432/advice?sslmode=disable")

	if err != nil {
		fmt.Println(err)
	}

	createTables(DB)

	return DB

}

func createTables(db *sql.DB) {

	query := `
		CREATE TABLE IF NOT EXISTS advice (
    		id SERIAL PRIMARY KEY,
    		name VARCHAR(255) NOT NULL
    	)`

	fmt.Println(db.Driver())

	_, err := db.Exec(query)

	if err != nil {
		panic(err)
	}

}
