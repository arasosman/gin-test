package repositories

import (
	"gin-test/inputs"
	"gin-test/models"
)

func FindBooks() []models.Book {
	var books []models.Book
	models.DB.Find(&books)
	return books
}

func FindBook(bookId int) (models.Book, error) {
	var book models.Book

	if err := models.DB.Where("id = ?", bookId).First(&book).Error; err != nil {
		return models.Book{}, err
	}
	return book, nil
}

func CreateBook(input inputs.CreateBookInput) models.Book {
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)
	return book
}
