package application

import (
	"earthly/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (app *App) PostCompany(c *gin.Context) {

	var company model.Companies

	if err := c.ShouldBindJSON(&company); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	if err := app.DB.Create(&company).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, company)

}

func (app *App) GetCompany(c *gin.Context) {
	var company model.Companies

	name := c.Param("company")

	if err := app.DB.Where("name= ?", name).First(&company).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Failed": "Company not found"})
		return
	}

	c.JSON(http.StatusOK, company)
}

func (app *App) UpdateCompany(c *gin.Context) {
	var company model.Companies

	name := c.Param("company")

	var err error

	if err = app.DB.Where("name = ?", name).First(&company).Error; err != nil {

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

	if err = app.DB.Model(&company).Updates(model.Companies{
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

func (app *App) DeleteCompany(ctx *gin.Context) {

	var company model.Companies

	name := ctx.Param("company")

	var err error

	if err = app.DB.Where("name = ?", name).First(&company).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "company not found!"})
		return
	}

	if err = app.DB.Delete(&company).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Company Deleted"})
}
