package handler

import (
	"database/sql"
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"net/http"
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

func render(c *gin.Context, status int, template templ.Component) error {

	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}

func (u *UsersHandler) Home(c *gin.Context) {
	err := render(c, 200, view.Hello("Mirodil"))

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
