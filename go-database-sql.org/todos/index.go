package todos

import "database/sql"

var db *sql.DB

type Todo struct {
	Id    int
	Title string
}

func SetDB(_db *sql.DB) {
	db = _db
}
