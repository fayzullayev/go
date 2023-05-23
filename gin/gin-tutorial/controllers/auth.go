package controllers

import (
	"gin-tutorial/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := &models.User{
		Username: input.Username,
		Password: input.Password,
	}

	saveUser, err := user.SaveUser()

	if err != nil {
		log.Println("Error:", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": saveUser,
	})
}
