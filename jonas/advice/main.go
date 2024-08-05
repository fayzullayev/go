package main

import (
	"advice/db"
	"advice/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db.InitDB()

	router.Use(cors.Default())

	router.GET("/advices", handlers.GetAdvices)

	err := router.Run(":8888")
	if err != nil {
		panic("error starting server")
	}

}
