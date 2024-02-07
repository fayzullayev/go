package event

import (
	"database/sql"
	"fmt"
	"rest/db"
	"time"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime"`
	UserID      int       `json:"userID"`
}

func (e *Event) Save() error {
	query := `
		INSERT INTO events (name, description, location, dateTime, user_id) 
		VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	e.ID = id

	//events = append(events, *e)
	return nil
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`

	var events []Event

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var event Event

		err = rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (e *Event) GetById() error {
	query := "SELECT * FROM events WHERE id = ?"

	row := db.DB.QueryRow(query, e.ID)

	err := row.Err()
	if err != nil {
		fmt.Println(11111)

		return err
	}

	if err = row.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID); err != nil {
		fmt.Println(222222)
		return err
	}

	return nil
}

func (e *Event) Update() error {
	query := `UPDATE events 
			  SET name = ?, description = ?,  dateTime = ?
			  WHERE id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer func(stmt *sql.Stmt) {
		err = stmt.Close()
		if err != nil {

		}
	}(stmt)

	_, err = stmt.Exec(e.Name, e.Description, e.DateTime, e.ID)

	if err != nil {
		return err
	}

	return nil

}
