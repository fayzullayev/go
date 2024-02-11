package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"rest/db"
	"rest/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	err := server.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal("server.SetTrustedProxies", err)
	}

	routes.RegisterRoutes(server)

	err = server.Run(":8888")
	if err != nil {
		log.Fatal("server.Run", err)
	}
}

//178-179-180-181-182
