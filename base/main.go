package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

type Todo struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("DATABEUSER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"))
}

func main() {
	db := initDB()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("Can not close the database connection")
		}
	}(db)

	err := db.Ping()

	if err != nil {
		log.Fatalf("Connection error(Ping) %v", err)
	}

}

func initDB() *sql.DB {
	db, err := sql.Open("pgx", DSN())

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
