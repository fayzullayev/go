package handler

import "database/sql"

type Handlers struct {
	DB           *sql.DB
	UsersHandler *UsersHandler
	TasksHandler *TasksHandler
}

func NewHandlers(db *sql.DB) *Handlers {
	return &Handlers{
		DB:           db,
		UsersHandler: NewUsersHandler(db),
		TasksHandler: NewTasksHandler(db),
	}
}
