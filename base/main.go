package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
)

type Todo struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func main() {
	fmt.Println("Hello World")

	db, err := initDB()
	defer db.Close()

	if err != nil {
		fmt.Println(1, err)

		return
	}

	err = db.Ping()

	if err != nil {
		fmt.Println(2, err)
		return
	}

	var todos []Todo

	rows, err := db.Query("SELECT * from todos")

	if err != nil {
		log.Fatal(err)
		return
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.Id, &todo.Title)

		if err != nil {
			log.Fatal(err)
		}

		todos = append(todos, todo)
	}

	err = rows.Err()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("todos", todos)
}

func initDB() (*sql.DB, error) {
	const (
		user     = "mirodil"
		password = "123456789"
		host     = "localhost"
		port     = 5432
		dbname   = "todos"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("pgx", psqlInfo)

	if err != nil {
		log.Fatalf("Error %s", err)
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		log.Fatalf("Error%s", err)
		return nil, err
	}

	return db, nil

}
