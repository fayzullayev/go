package controllers

import (
	"earthly/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func PostCompany(db *gorm.DB) gin.HandlerFunc {

	return func(c *gin.Context) {

		var company model.Companies

		if err := c.ShouldBindJSON(&company); err != nil {

			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := db.Create(&company).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, company)
	}

}

func GetCompany(db *gorm.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		var company model.Companies

		name := c.Param("company")

		if err := db.Where("name= ?", name).First(&company).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"Failed": "Company not found"})
			return
		}

		c.JSON(http.StatusOK, company)
	}
}

func UpdateCompany(db *gorm.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		var company model.Companies

		name := c.Param("company")

		var err error

		if err = db.Where("name = ?", name).First(&company).Error; err != nil {

			c.JSON(http.StatusNotFound, gin.H{
				"error": "Company doesn't exist",
			})

			return
		}

		if err = c.ShouldBindJSON(&company); err != nil {

			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err = db.Model(&company).Updates(model.Companies{
			Name:    company.Name,
			Created: company.Created,
			Product: company.Product,
		}).Error; err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, company)
	}

}

func DeleteCompany(db *gorm.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var company model.Companies

		name := ctx.Param("company")

		var err error

		if err = db.Where("name = ?", name).First(&company).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "company not found!"})
			return
		}

		if err = db.Delete(&company).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError,
				gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Company Deleted"})
	}

}
