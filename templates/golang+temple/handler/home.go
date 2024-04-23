package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"tempaltes-tutorial/model"
	"tempaltes-tutorial/view"
)

type HomeHandler struct {
	db *sql.DB
}

func NewHomeHandler(db *sql.DB) *HomeHandler {
	return &HomeHandler{
		db: db,
	}
}

func (h *HomeHandler) HomePage(c *gin.Context) {
	user := model.NewUser("Mirodil3")

	err := render(c, 200, view.Hello("Hi Buddy2", user))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
}
