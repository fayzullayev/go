package handler

import (
	"database/sql"
)

type Handlers struct {
	HomeHandler  *HomeHandler
	UsersHandler *UsersHandler
	TasksHandler *TasksHandler
}

func NewHandlers(db *sql.DB) *Handlers {
	return &Handlers{
		HomeHandler:  NewHomeHandler(db),
		UsersHandler: NewUsersHandler(db),
		TasksHandler: NewTasksHandler(db),
	}
}
