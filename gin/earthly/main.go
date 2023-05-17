package main

import (
	"earthly/controllers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	router.POST("/company", controllers.PostCompany)

	log.Fatal(router.Run(":6666"))

}
