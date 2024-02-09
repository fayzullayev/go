package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {

	events := server.Group("/events")
	{
		events.GET("/", getEvents)
		events.GET("/:eventId", getEvent)
		events.POST("/", createEvents)
		events.PUT("/:eventId", updateEvent)
		events.DELETE("/:eventId", deleteEvent)
	}

	auth := server.Group("/auth")
	{
		auth.POST("signup", signUp)
	}

}
