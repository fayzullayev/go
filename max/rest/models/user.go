package models

import (
	"database/sql"
	"rest/db"
	"rest/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users(email, password) VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)
	defer func(stmt *sql.Stmt) {
		err = stmt.Close()
		if err != nil {

		}
	}(stmt)

	if err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId

	return err

}

func (u *User) FindByEmail() error {
	query := `SELECT * FROM users WHERE email = ?`

	stmt, err := db.DB.Prepare(query)

	defer func(stmt *sql.Stmt) {
		err = stmt.Close()
		if err != nil {

		}
	}(stmt)

	if err != nil {
		return err
	}

	row := stmt.QueryRow(u.Email)

	err = row.Scan(&u.ID, &u.Email, &u.Password)
	if err != nil {
		return err
	}

	return nil

}
