package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"rest/utils"
	"strings"
)

func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized1",
		})
		return
	}

	fmt.Println(strings.Split(token, " ")[1])

	userId, err := utils.VerifyToken(strings.Split(token, " ")[1])

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized2",
		})
		return
	}

	c.Set("userId", userId)

	c.Next()
}
