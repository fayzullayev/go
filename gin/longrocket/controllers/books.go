package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"longrocket/models"
)

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// FindBooks GET /books
// Get all books
func FindBooks(c *gin.Context) {
	var books []models.Book

	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func CreateBook(c *gin.Context) {
	// Validate input
	var input CreateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}
