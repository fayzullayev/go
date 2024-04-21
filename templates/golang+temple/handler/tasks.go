package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TasksHandler struct {
	DB *sql.DB
}

func NewTasksHandler(db *sql.DB) *TasksHandler {
	return &TasksHandler{
		DB: db,
	}
}

func (t *TasksHandler) GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"tasks": "all tasks",
	})
}
