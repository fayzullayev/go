package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	var r *gin.Engine

	r = gin.Default()

	r.GET("/books", getAllBooks)
	r.POST("/books", createBook)
	r.DELETE("/books/:id", deleteBook)

	log.Fatal(r.Run(":8008"))
}

type Book struct {
	ID     string `json:"ids"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: "1", Title: "Harry Potter", Author: "J. K. Rowling"},
	{ID: "2", Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
	{ID: "3", Title: "The Wizard of Oz", Author: "L. Frank Baum"},
}

func getAllBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var book Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Wrong fields"})
		return
	}

	books = append(books, book)

	c.JSON(http.StatusCreated, book)
}

func deleteBook(c *gin.Context) {
	id := c.Param("id")

	for i, a := range books {
		if a.ID == id {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}

	c.Status(http.StatusNoContent)
}
