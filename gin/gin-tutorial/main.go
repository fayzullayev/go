package main

import (
	"gin-tutorial/controllers"
	"gin-tutorial/models"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	models.ConnectDataBase()

	router := gin.Default()

	public := router.Group("/api")

	public.POST("/register", controllers.Register)

	log.Fatal(router.Run(":9999"))

}
