package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"rest/db"
	"rest/models/event"
	"time"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	db.InitDB()
	server := gin.Default()

	err := server.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal("server.SetTrustedProxies", err)
	}

	server.GET("/events", getEvents)
	server.POST("/events", createEvents)

	err = server.Run(":8888")
	if err != nil {
		log.Fatal("server.Run", err)
	}
}

func getEvents(c *gin.Context) {
	events, err := event.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch events: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, events)
}

func createEvents(c *gin.Context) {
	var myEvent event.Event

	if err := c.ShouldBindJSON(&myEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something went wrong:" + err.Error(),
		})
		return
	}

	myEvent.ID = 1
	myEvent.UserID = 1
	myEvent.DateTime = time.Now()

	err := myEvent.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong: " + err.Error(),
		})
		return
	}

	//events := event.GetAllEvents()
	c.JSON(http.StatusCreated, myEvent)
}
