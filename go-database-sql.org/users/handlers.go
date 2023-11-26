package users

import (
	"database/sql"
	"log"
)

func (u *User) GetUsersById() error {
	stmt, err := db.Prepare("SELECT id, firstName, lastName, age, job, phoneNumber FROM users WHERE id = ?")

	if err != nil {
		return err
	}

	defer func(stmt *sql.Stmt) {
		err = stmt.Close()
		if err != nil {
			log.Fatal("Closing error")
		}
	}(stmt)

	err = stmt.QueryRow(u.Id).Scan(&u.Id, &u.FirstName, &u.LastName, &u.Age, &u.Job, &u.PhoneNumber)

	if err != nil {
		return err
	}

	return nil
}

func GetUsers() ([]User, error) {
	var users []User
	rows, err := db.Query("SELECT * FROM  users LIMIT  5")

	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			log.Fatal("Closing error")
		}
	}(rows)

	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Age, &user.Job, &user.PhoneNumber)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
