package models

import (
	"database/sql"
	"errors"
	"log"
	"project-managememnt/db"
)

type Project struct {
	ID          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	DueDate     string `json:"dueDate" binding:"required"`
}

func (p *Project) Save() (int, error) {
	query := "INSERT INTO projects (title, description, duedate) VALUES ($1, $2, $3) RETURNING id"

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

	err = stmt.QueryRow(p.Title, p.Description, p.DueDate).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (p *Project) GetByID() error {
	query := "SELECT title, description, duedate FROM projects WHERE id = $1"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer func(stmt *sql.Stmt) {
		err = stmt.Close()
		if err != nil {
			log.Fatal("Error closing stmt")
		}
	}(stmt)

	err = stmt.QueryRow(p.ID).Scan(&p.Title, &p.Description, &p.DueDate)
	if err != nil {
		return err
	}
	return nil
}

func GetProjects() ([]Project, error) {
	query := "SELECT * FROM projects"

	var projects []Project = []Project{}

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

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var project Project

		err = rows.Scan(&project.ID, &project.Title, &project.Description, &project.DueDate)
		if err != nil {
			return nil, err
		}

		projects = append(projects, project)
	}

	err = rows.Close()
	if err != nil {
		return nil, err
	}

	return projects, nil

}

func (p *Project) Delete() error {
	query := "DELETE FROM projects WHERE id = $1"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err = stmt.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(stmt)

	result, err := stmt.Exec(p.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()

	if err != nil {
		return err
	} else {
		if rows != 0 {
			return nil
		}
	}

	return errors.New("project not found")
}
