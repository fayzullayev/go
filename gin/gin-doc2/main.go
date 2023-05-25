package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	var router *gin.Engine

	router = gin.Default()

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, map[string]interface{}{
			"Hello": 8798798,
		})
	})

	log.Fatal(router.Run(":9999"))

}
