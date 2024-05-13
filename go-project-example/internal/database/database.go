package database

import (
	"database/sql"
	"fmt"
)
import _ "github.com/lib/pq"

func Init() (*sql.DB, error) {
	const driver string = "postgres"

	const conString = "postgres://mirodil:12345@localhost:5432/test?sslmode=disable"

	db, err := sql.Open(driver, conString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	err = createTables(db)
	if err != nil {
		return nil, err
	}

	fmt.Println("Database successfully initialized")

	return db, nil
}

func createTables(db *sql.DB) error {
	var err error

	const userTableCreate = `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY, 
			username VARCHAR (50) UNIQUE NOT NULL, 
			password VARCHAR (50) NOT NULL, 
			email VARCHAR (255) UNIQUE NOT NULL, 
			created_at TIMESTAMP NOT NULL, 
			last_login TIMESTAMP
		)`

	const eventsTableCreate = `
		CREATE TABLE IF NOT EXISTS events (
			id SERIAL PRIMARY KEY, 
			title VARCHAR (255) NOT NULL, 
			description VARCHAR (255), 
			created_at TIMESTAMP NOT NULL, 
			last_login TIMESTAMP
		)`

	_, err = db.Exec(userTableCreate)

	if err != nil {
		return err
	}

	_, err = db.Exec(eventsTableCreate)

	if err != nil {
		return err
	}

	fmt.Println("Tables were created")

	return nil
}
