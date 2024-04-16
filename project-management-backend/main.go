package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"project-managememnt/routes"
)

func main() {

	initDB()

	router := gin.Default()

	err := router.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal("server.SetTrustedProxies", err)
	}

	routes.RegisterRoutes(router)

	err = router.Run(":8088")
	if err != nil {
		log.Fatal(err)
	}
}
