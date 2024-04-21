package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"project-managememnt/db"
	"project-managememnt/routes"
)

func main() {

	db.InitDB()

	router := gin.Default()

	err := router.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal("server.SetTrustedProxies", err)
	}

	router.Use(cors.Default())

	routes.RegisterRoutes(router)

	err = router.Run(":8088")
	if err != nil {
		log.Fatal(err)
	}
}
