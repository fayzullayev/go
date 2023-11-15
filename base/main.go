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

type User struct {
	Id          int    `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Age         int    `json:"age"`
	Job         string `json:"job"`
	PhoneNumber string `json:"phoneNumber"`
}

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {

	db := initDB()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("Can not close the database connection")
		}
	}(db)

	rows, err := db.Query("SELECT * FROM users")

	if err != nil {
		log.Fatalf("1. Database error %v", err)
	}

	var todos []User

	for rows.Next() {
		var todo User
		err = rows.Scan(&todo.Id, &todo.FirstName, &todo.LastName, &todo.Age, &todo.Job, &todo.PhoneNumber)

		if err != nil {
			log.Fatalf("2. Database error %v", err)
		}
		todos = append(todos, todo)
	}

	for i, v := range todos {
		fmt.Printf("%d. %+v\n", i, v)
	}
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

func albumsByArtist(name string, db *sql.DB) ([]Album, error) {
	// An albums slice to hold data from returned rows.
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}
