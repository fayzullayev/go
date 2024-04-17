package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := router.Group("/api/v1")
	projects := v1.Group("/projects")

	{
		projects.GET("", GetProjectsHandler)
		projects.GET("/:id", GetProjectHandler)
		projects.DELETE("/:id", DeleteProjectHandler)
		projects.POST("", CreateProjectHandler)
	}

	tasks := v1.Group("/tasks")

	{
		tasks.GET("/:id", GetTasksHandler)
		tasks.POST("", CreateTasksHandler)
	}

}
