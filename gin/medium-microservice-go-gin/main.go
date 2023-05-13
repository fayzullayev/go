package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"time"
)

type EmailJob struct {
	Sender  string `json:"sender" binding:"required"`
	JobId   int    `json:"job_id" binding:"gte=0"`
	Subject string `json:"subject" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Body    string `json:"body" binding:"required"`
}

func main() {

	router := gin.Default()

	router.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "Email Server Running"})
	})

	router.POST("/send-email", func(c *gin.Context) {
		var emailHandler EmailJob

		if err := c.ShouldBindJSON(&emailHandler); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input!", "er": err})
			return
		}

		log.Printf("Email Service: creating new Email job #%v...", emailHandler.Sender)
		rand.Seed(time.Now().UnixNano())
		emailHandler.JobId = rand.Intn(1000)
		log.Printf("EmailService: created print job #%v", emailHandler.JobId)

		//send email here
		// sendEmail(emailHandler.Subject, emailHandler.Body+" has been shipped")
		//
		c.JSON(200, emailHandler)
	})

	log.Fatal(router.Run(":9000"))
}
