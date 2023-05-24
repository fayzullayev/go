package main

import (
	"gin-tutorial/controllers"
	"gin-tutorial/middlewares"
	"gin-tutorial/models"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	models.ConnectDataBase()

	router := gin.Default()

	public := router.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := router.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())

	protected.GET("/user", controllers.CurrentUser)

	log.Fatal(router.Run(":9999"))

}
