package controllers

import (
	"gin-tutorial/models"
	"gin-tutorial/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var input UserInput

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
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": saveUser,
	})
}

func Login(c *gin.Context) {
	var input UserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	user := models.User{
		Username: input.Username,
		Password: input.Password,
	}

	token, err := models.LoginCheck(user.Username, user.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func CurrentUser(c *gin.Context) {

	userId, err := utils.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetUserByID(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}
