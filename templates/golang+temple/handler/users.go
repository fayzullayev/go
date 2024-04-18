package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UsersHandler struct {
	DB *sql.DB
}

func NewUsersHandler(db *sql.DB) *UsersHandler {
	return &UsersHandler{
		DB: db,
	}
}

func (u *UsersHandler) GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"users": "all users",
	})
}
