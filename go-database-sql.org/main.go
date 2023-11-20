package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	db, err := sql.Open("sqlite3", "./todos.db")

	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("--------------------------------------")
	fmt.Println("Successfully connected to the database")
	fmt.Println("--------------------------------------")

}
