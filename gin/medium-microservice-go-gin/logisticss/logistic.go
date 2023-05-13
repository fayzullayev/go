package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"log"
)

type EmailJob struct {
	Sender  string `json:"sender" binding:"required"`
	JobId   int    `json:"job_id" binding:"gte=0"`
	Subject string `json:"subject" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Body    string `json:"body" binding:"required"`
}

type LogisticsJob struct {
	Type  string `json:"type" binding:"required"`
	Email string `json:"email" binding:"required"`
	Item  string `json:"item" binding:"required"`
}

func createEmailJob(emailJob EmailJob) EmailJob {
	client := resty.New()

	var p EmailJob

	_, err := client.R().
		SetBody(emailJob).
		SetResult(&p).
		Post("http://localhost:9000/send-email")

	if err != nil {
		log.Println("EmailServer: unable to connect EmailService")
		return EmailJob{}
	}

	log.Printf("EmailServer: created email job #%v via EmailService", p.JobId)

	return p
}

func main() {
	router := gin.Default()

	router.GET("/status", func(c *gin.Context) {
		c.JSON(200, "Logistics Server Running")
	})

	router.POST("/process", func(c *gin.Context) {

		var logistics LogisticsJob

		if err := c.ShouldBindJSON(&logistics); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input!", "er": err.Error()})
			return
		}

		//processing
		emailData := EmailJob{
			Sender:  "abc@logistics.com",
			Subject: "Item" + logistics.Type,
			Email:   logistics.Email,
			Body:    logistics.Item,
		}
		//save this data in database

		//calling email server

		emailData = createEmailJob(emailData)

		c.JSON(200, emailData)
	})

	log.Fatal(router.Run(":9001"))
}
