package models

import (
	"database/sql"
	"fmt"
	"log"
	"project-managememnt/db"
)

type Task struct {
	Id        int    `json:"id"`
	Title     string `json:"title" binding:"required"`
	ProjectID int    `json:"projectID,omitempty" binding:"required"`
}

func (t *Task) Save() (int, error) {
	query := "INSERT INTO tasks (title, project_id) VALUES ($1, $2) RETURNING id"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return 0, err
	}

	defer func(stmt *sql.Stmt) {
		err = stmt.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(stmt)

	var id int

	if err = stmt.QueryRow(t.Title, t.ProjectID).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func GetTasks(projectId int) ([]Task, error) {
	query := "SELECT id, title FROM tasks where project_id = $1"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer func(stmt *sql.Stmt) {
		err = stmt.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(stmt)

	rows, err := stmt.Query(projectId)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	var tasks []Task = []Task{}

	for rows.Next() {
		var task Task

		err = rows.Scan(&task.Id, &task.Title)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	fmt.Println("tasks:", tasks)

	return tasks, nil
}
