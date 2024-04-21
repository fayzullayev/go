package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {

	router.GET("/ping", func(c *gin.Context) {

		c.BindJSON()
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := router.Group("/api/v1")
	projects := v1.Group("/projects")
	{
		projects.GET("/")
	}

}
