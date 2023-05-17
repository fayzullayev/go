package controllers

import (
	"earthly/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func PostCompany(c *gin.Context) {
	var company model.Companies

	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	db, err := model.DBConnection()

	if err != nil {
		log.Println(err)
		return
	}

	if err = db.Create(&company).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, company)
}
