package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		clientIP := c.ClientIP()

		now := time.Now()

		log.Printf("[%s] %s %s %s", now.Format(time.RFC3339), c.Request.Method, c.Request.URL.Path, clientIP)

		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.Use(Logger())
	router.Run(":8087")
}
