package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"sqltutorial/users"
)

func main() {
	db, err := setupDB()

	if err != nil {
		log.Fatal(err)
	}

	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	fmt.Println("--------------------------------------")
	fmt.Println("Successfully connected to the database")
	fmt.Println("--------------------------------------")

	usersList, err := users.GetUsers()

	if err != nil {
		fmt.Printf("Something went wrong:\n%v", err.Error())
		return
	}

	for index, value := range usersList {
		fmt.Printf("%d. %+v\n", index, value)
	}

	fmt.Println("--------------------------------------")

	user := users.User{Id: 25}

	err = user.GetUsersById()

	if err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Printf("User: %+v", user)

}

func setupDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./todos.db")

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	users.SetDB(db)

	return db, nil
}
