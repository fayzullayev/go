package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *sql.DB

func InitDB() *sql.DB {
	var err error

	DB, err = sql.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatal(err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to the database")
	createTables()
	return DB
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)`

	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			user_id INTEGER,
			FOREIGN KEY(user_id) REFERENCES users(id)
		)`

	if DB != nil {
		_, err := DB.Exec(createUsersTable)
		if err != nil {
			log.Fatal("DB.Exec", err)
		}

		_, err = DB.Exec(createEventsTable)
		if err != nil {
			log.Fatal("DB.Exec", err)
		}

		fmt.Println("Events and Users tables were created successfully")
	}

}
