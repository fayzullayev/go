package main

import (
	"github.com/gin-gonic/gin"
	"longrocket/controllers"
	"longrocket/models"
)

func main() {

	//var router *gin.Engine

	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/books", controllers.FindBooks)
	router.POST("/books", controllers.CreateBook)
	router.GET("/books/:id", controllers.FindBook)
	router.PATCH("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	err := router.Run(":3004")

	if err != nil {
		return
	}

}
