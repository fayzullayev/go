package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"tempaltes-tutorial/model"
	"tempaltes-tutorial/view"
)

type UsersHandler struct {
	DB *sql.DB
}

func NewUsersHandler(db *sql.DB) *UsersHandler {
	return &UsersHandler{
		DB: db,
	}
}

func (u *UsersHandler) Home(c *gin.Context) {

	c.Status(200)

	err := view.Hello("John doe", model.User{
		Name: "John doe",
	}).Render(c.Request.Context(), c.Writer)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "error",
		})
		return
	}
}

func (u *UsersHandler) GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"users": "all users",
	})
}
