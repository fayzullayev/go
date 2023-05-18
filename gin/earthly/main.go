package main

import (
	"earthly/application"
	"earthly/model"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	db, err := model.DBConnection()

	if err != nil {
		panic("an error occurred during the connection to the database")
	}

	router := gin.Default()

	app := application.App{
		DB: db,
	}

	router.GET("api/v1/:company", app.GetCompany)
	router.POST("/company", app.PostCompany)
	router.PUT("api/v1/:company", app.UpdateCompany)
	router.DELETE("api/v1/:company", app.DeleteCompany)

	log.Fatal(router.Run(":6666"))

}
