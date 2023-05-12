package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type Handler struct {
	db *gorm.DB
}

func newHandler(db *gorm.DB) *Handler {
	return &Handler{db}
}

var db *gorm.DB

func main() {
	var err error

	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&Book{})

	if err != nil {
		panic("failed to connect migrate")
	}

	handler := newHandler(db)

	r := gin.Default()

	r.GET("/books", handler.getAllBooks)
	r.POST("/books", handler.createBook)
	r.DELETE("/books/:id", handler.deleteBook)

	log.Fatal(r.Run(":8008"))
}

type Book struct {
	ID     string `json:"ids" `
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (h *Handler) getAllBooks(c *gin.Context) {
	var books []Book

	if result := h.db.Find(&books); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func (h *Handler) createBook(c *gin.Context) {
	var book Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := h.db.Create(&book); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
	}

	c.JSON(http.StatusCreated, book)
}

func (h *Handler) deleteBook(c *gin.Context) {
	id := c.Param("id")

	if result := h.db.Delete(&Book{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
