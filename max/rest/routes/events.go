package routes

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"rest/models"
	"strconv"
	"time"
)

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch events: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, events)
}
func getEvent(c *gin.Context) {
	eventId, err := strconv.Atoi(c.Param("eventId"))

	//fmt.Printf("%+v", c.Params[0])

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found: " + err.Error(),
		})
		return
	}

	var e models.Event
	e.ID = int64(eventId)

	err = e.GetById()
	if err != nil {

		if errors.Is(sql.ErrNoRows, err) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "user not found.....: ",
				"error":   fmt.Sprintf("User with id:%v does not exist", eventId),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong...",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successful",
		"data":    e,
	})
}
func createEvents(c *gin.Context) {
	userId := c.GetInt64("userId")

	var myEvent models.Event

	if err := c.ShouldBindJSON(&myEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something went wrong:" + err.Error(),
		})
		return
	}

	myEvent.UserID = userId
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
func updateEvent(c *gin.Context) {
	eventId, err := strconv.Atoi(c.Param("eventId"))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "event not found-1",
		})

		return
	}

	myEvent := models.Event{ID: int64(eventId)}

	err = myEvent.GetById()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "event not found-2",
		})

		return
	}

	userId := c.GetInt64("userId")

	if userId != myEvent.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not allowed",
		})

		return
	}

	err = c.ShouldBindJSON(&myEvent)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "event not found-3" + err.Error(),
		})
		return
	}

	err = myEvent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "event not found-4" + err.Error(),
		})
		return
	}

	err = myEvent.GetById()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "event not found-2",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Ok",
		"data":    myEvent,
	})
}
func deleteEvent(c *gin.Context) {
	eventId := c.Param("eventId")

	id, err := strconv.Atoi(eventId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "event not found",
		})
		return
	}

	myEvent := models.Event{ID: int64(id)}

	err = myEvent.GetById()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "event not found",
		})
		return
	}

	userId := c.GetInt64("userId")

	if userId != myEvent.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not allowed",
		})

		return
	}

	err = myEvent.Delete()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successful",
	})
}
