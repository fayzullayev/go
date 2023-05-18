package main

import (
	"earthly/controllers"
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

	router.GET("api/v1/:company", controllers.GetCompany(db))
	router.POST("/company", controllers.PostCompany(db))
	router.PUT("api/v1/:company", controllers.UpdateCompany(db))
	router.DELETE("api/v1/:company", controllers.DeleteCompany(db))

	log.Fatal(router.Run(":6666"))

}
