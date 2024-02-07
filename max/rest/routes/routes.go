package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:eventId", getEvent)
	server.POST("/events", createEvents)
	server.PUT("/events/:eventId", updateEvent)
}
