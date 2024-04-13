package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

func initDB() {
	connStr := "postgres://postgres:12345@localhost/projects_management?sslmode=disable"

	var err error

	DB, err = sql.Open("postgres", connStr)

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
}

func createTables() {

	createProjectsTable := `
		CREATE TABLE IF NOT EXISTS projects (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			description VARCHAR(255) NOT NULL,
			dueDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`

	createTasksTable := `
		CREATE TABLE IF NOT EXISTS tasks (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			project_id INTEGER NOT NULL,
			FOREIGN KEY(project_id) REFERENCES projects(id)
		)`

	if DB != nil {
		_, err := DB.Exec(createProjectsTable)
		_, err = DB.Exec(createTasksTable)

		if err != nil {
			log.Fatal("DB.Exec", err)
		}

		fmt.Println("Projects and Tasks tables were created successfully")
	}
}
