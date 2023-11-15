package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

type Todo struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func main() {

	db := initDB()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("Can not close the database connection")
		}
	}(db)

	rows, err := db.Query("SELECT * FROM todos")

	if err != nil {
		log.Fatalf("1. Database error %v", err)
	}

	var todos []Todo

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.Id, &todo.Title)

		if err != nil {
			log.Fatalf("2. Database error %v", err)
		}
		todos = append(todos, todo)
	}

	fmt.Printf("%+v", todos)
}

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./todos.db")

	if err != nil {
		log.Fatalf("Connection error %v", err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatalf("Connection error(Ping) %v", err)
	}

	fmt.Println("Successfully connected")

	return db
}

func DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("DATABASE_USER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"))
}
