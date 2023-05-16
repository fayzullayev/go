package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to building RESTful api",
		})

	})

	log.Fatal(r.Run(":5005"))
}
