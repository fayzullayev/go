package todos

import (
	"database/sql"
	"fmt"
)

func (t *Todo) GetTodoById() error {
	if err := db.QueryRow("SELECT * FROM todos WHERE id = ?", t.Id).Scan(&t.Id, &t.Title); err != nil {
		return err
	}
	return nil
}

func GetTodos() ([]Todo, error) {
	rows, err := db.Query("SELECT * FROM todos")

	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			fmt.Println("Error")
		}
	}(rows)

	var todos []Todo

	for rows.Next() {
		var todo Todo
		if err = rows.Scan(&todo.Id, &todo.Title); err != nil {
			fmt.Println("Error")
		}
		todos = append(todos, todo)
	}

	return todos, nil
}
