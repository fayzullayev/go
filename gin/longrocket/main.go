package main

import (
	"github.com/gin-gonic/gin"
	"longrocket/controllers"
	"longrocket/models"
)

func main() {

	var router *gin.Engine

	router = gin.Default()

	models.ConnectDatabase()

	router.GET("/books", controllers.FindBooks)

	err := router.Run(":3004")
	
	if err != nil {
		return
	}

}
