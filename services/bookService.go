package services

import (
	"gin-test/inputs"
	"gin-test/models"
	"gin-test/repositories"
	"github.com/gin-gonic/gin"
	"strconv"
)

func FindBooks(c *gin.Context) []models.Book {
	return repositories.FindBooks()
}

func FindBook(c *gin.Context) (models.Book, error) {
	bookId, _ := strconv.Atoi(c.Param("id"))
	book, err := repositories.FindBook(bookId)
	if err != nil {
		return book, err
	}
	return book, nil
}

func CreateBook(c *gin.Context) (models.Book, error) {
	// Validate input
	var input inputs.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		return models.Book{}, err
	}
	return repositories.CreateBook(input), nil
}
