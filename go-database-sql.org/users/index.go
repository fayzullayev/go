package users

import (
	"database/sql"
)

var db *sql.DB

type User struct {
	Id          int
	FirstName   string
	LastName    string
	Age         int
	Job         string
	PhoneNumber string
}

func SetDB(_db *sql.DB) {
	db = _db
}
