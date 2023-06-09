package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"longrocket/models"
	"net/http"
)

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// FindBooks GET /books
// Get all books
func FindBooks(c *gin.Context) {
	var books []models.Book

	models.DB.Find(&books)

	if len(books) == 0 {
		c.JSON(http.StatusOK, gin.H{"data": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": books})

}

func FindBook(c *gin.Context) {
	var book models.Book
	fmt.Println(book)
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	fmt.Println(book)
	c.JSON(http.StatusOK, gin.H{"data": book})

}

func CreateBook(c *gin.Context) {
	// Validate input
	var input CreateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{Title: input.Title, Author: input.Author}

	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(c *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})

	//if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	//	return
	//}
}

func DeleteBook(c *gin.Context) {
	var book models.Book
	fmt.Println("book", book)
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	fmt.Println("book", book)
	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
