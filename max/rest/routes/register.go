package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"rest/models"
	"strconv"
)

func registerForEvent(c *gin.Context) {
	eventId := c.Param("eventId")
	userId := c.GetInt64("userId")

	id, err := strconv.Atoi(eventId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "event not found",
		})
		return
	}

	event := models.Event{ID: int64(id)}

	err = event.GetById()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "event not found",
		})
		return
	}

	fmt.Printf("%+v", event)

	err = event.Register(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not register",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "registered1",
	})
}

func cancelRegistration(c *gin.Context) {
	eventId := c.Param("eventId")
	userId := c.GetInt64("userId")

	id, err := strconv.Atoi(eventId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "event not found",
		})
		return
	}

	event := models.Event{ID: int64(id)}

	err = event.CancelRegistration(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not cancel",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "cancelled",
	})

}
