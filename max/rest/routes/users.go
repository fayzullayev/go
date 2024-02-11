package routes

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"rest/models"
	"rest/utils"
)

func signUp(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message1": err.Error(),
		})
		return
	}

	err := user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created",
	})

}

func signIn(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something went wrong1",
		})

		return
	}

	initialPassword := user.Password

	err := user.FindByEmail()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "email or password incorrect",
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong2",
			"error":   err.Error(),
		})

		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(initialPassword))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "email or password incorrect",
		})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong2",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"token":   token,
		"data":    user,
	})

	return
}
