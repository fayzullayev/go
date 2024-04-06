package main

import (
	"database/sql"
	"errors"
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

func (u *User) Create(db *sql.DB) (int64, error) {
	result, err := db.Exec("INSERT INTO users (firstName, lastName, age, job, phoneNumber) VALUES (?, ?, ?, ?, ?)", u.FirstName, u.LastName, u.Age, u.Job, u.PhoneNumber)

	if err != nil {
		return 0, fmt.Errorf("create user: %v", err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("create user: %v", err)
	}

	num, err := result.RowsAffected()

	if err != nil {
		return 0, fmt.Errorf("create user: %v", err)
	}

	fmt.Println("RowsAffected", num)

	return id, nil
}

func (u *User) GetUserById(db *sql.DB, userId int64) (User, error) {
	//var user User
	//err := db.QueryRow("SELECT * FROM users WHERE id = ?", userId).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Age, &user.Job, &user.PhoneNumber)
	//
	//if err != nil {
	//	return user, err
	//}
	//return user, nil
	//
	var user User

	if err := db.QueryRow("SELECT * FROM users WHERE id = ?", userId).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Age, &user.Job, &user.PhoneNumber); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, fmt.Errorf("albumsById %d: no such album", userId)
		}
		return user, fmt.Errorf("albumsById %d: %v", userId, err)
	}
	return user, nil

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

	user := User{
		FirstName:   "Odil",
		LastName:    "Fayz",
		Age:         29,
		Job:         "Dentist",
		PhoneNumber: "+998909711001",
	}

	num, err := user.Create(db)

	if err != nil {
		fmt.Println("errerrerr", err)
	}

	fmt.Println("user created id:", num)

	user, err = user.GetUserById(db, num)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("%+v", user)

	//users, err := getUsers(db, "Journalist")
	//
	//if err != nil {
	//	fmt.Println("err", err)
	//	return
	//}
	//
	//for i, v := range users {
	//	fmt.Printf("%d. %+v\n", i, v)
	//}
	//
	//fmt.Println("-----------------------------------")
	//
	//user, err := getUserByID(db, 18)
	//
	//if err != nil {
	//	fmt.Println("err", err)
	//	return
	//}
	//
	//fmt.Printf("User: %+v", user)

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

func getUsers(db *sql.DB, name string) ([]User, error) {
	var users []User

	rows, err := db.Query("SELECT * FROM users WHERE job = ?", name)
	if err != nil {
		return nil, fmt.Errorf("getUsers %q: %v", name, err)
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Age, &user.Job, &user.PhoneNumber); err != nil {
			return nil, fmt.Errorf("getUsers %q: %v", name, err)
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("getUsers %q: %v", name, err)
	}
	return users, nil
}

func getUserByID(db *sql.DB, id int64) (User, error) {
	// An album to hold data from the returned row.
	var user User

	if err := db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Age, &user.Job, &user.PhoneNumber); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, fmt.Errorf("albumsById %d: no such album", id)
		}
		return user, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return user, nil
}
