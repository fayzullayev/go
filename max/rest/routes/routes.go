package routes

import (
	"github.com/gin-gonic/gin"
	"rest/middlewares"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "works",
		})
	})

	events := server.Group("/events")

	{
		events.GET("/", getEvents)
		events.GET("/:eventId", getEvent)

		authEvents := events.Group("/")
		authEvents.Use(middlewares.Authenticate)
		{
			authEvents.POST("/", createEvents)
			authEvents.PUT("/:eventId", updateEvent)
			authEvents.DELETE("/:eventId", deleteEvent)
			authEvents.POST("/:eventId/register", registerForEvent)
			authEvents.DELETE("/:eventId/register", cancelRegistration)
		}
	}

	auth := server.Group("/auth")
	{
		auth.POST("/signup", signUp)
		auth.POST("/login", signIn)
	}

}
