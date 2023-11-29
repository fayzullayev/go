package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"sqltutorial/todos"
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
	//
	//usersList, err := users.GetUsers()
	//
	//if err != nil {
	//	fmt.Printf("Something went wrong :\n%v", err.Error())
	//	return
	//}
	//
	//for index, value := range usersList {
	//	fmt.Printf("%d. %+v\n", index, value)
	//}
	//
	//fmt.Println("--------------------------------------")
	//
	//user := users.User{Id: 25}
	//
	//err = user.GetUsersById()
	//
	//if err != nil {
	//	fmt.Println("error", err)
	//	return
	//}
	//
	//fmt.Printf("User: %+v\n", user)
	//
	//fmt.Println("---------------Todos-----------------")
	//
	//todo := todos.Todo{Id: 4}
	//
	//err = todo.GetTodoById()
	//
	//if err != nil {
	//	log.Fatal("User does not exist")
	//}
	//
	//fmt.Printf("Todo: %+v\n", todo)
	//
	//todosList, err := todos.GetTodos()
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for index, value := range todosList {
	//	fmt.Printf("%d. %+v\n", index, value)
	//}

	user := users.User{
		FirstName:   "Zaynab",
		LastName:    "To'qinova",
		Age:         3,
		Job:         "Baby",
		PhoneNumber: "+998900000000",
	}

	err = user.Save()

	if err != nil {
		fmt.Printf("Something went wrong :\n%v", err.Error())
		return
	}

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
	todos.SetDB(db)

	return db, nil
}
