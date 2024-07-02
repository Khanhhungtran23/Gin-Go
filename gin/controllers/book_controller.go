package controllers

import (
	"github.com/gin-gonic/gin"
	"myproject/config"
	"myproject/models"
	"net/http"
)

// CreateBook handles POST /books to create a new book
func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&book)
	c.JSON(http.StatusOK, book)
}

// GetBooks handles GET /books to retrieve a list of books
func GetBooks(c *gin.Context) {
	var books []models.Book
	config.DB.Find(&books)
	c.JSON(http.StatusOK, books)
}

// GetBook handles GET /books/:id to retrieve a single book by its ID
func GetBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// UpdateBook handles PUT /books/:id to update a book by its ID
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&book)
	c.JSON(http.StatusOK, book)
}

// DeleteBook handles DELETE /books/:id to delete a book by its ID
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Book{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
