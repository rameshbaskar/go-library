package services

import (
	"fmt"
	"go-library/database"
	"go-library/models"
)

type BookService struct {

}

func (bs BookService) FindOrCreateBook(
	isbn string,
	title string,
	author string,
) (models.Book, error) {
	existingBook, err := bs.FindByISBN(isbn)
	if err != nil {
		fmt.Println(err)
	}

	if existingBook.ISBN == isbn {
		fmt.Println("Book already exists. No need to create again.")
		return existingBook, nil
	}

	fmt.Println("Book does not exist. Need to create a new one.")
	newBook := models.Book{
		ISBN: isbn,
		Title: title,
		Author: author,
		CreatedBy: "system",
		UpdatedBy: "system",
	}
	db := database.GetDbConnection()
	createResult := db.Omit("id").Create(&newBook)
	db.Commit()
	return newBook, createResult.Error
}

func (bs BookService) FindByISBN(isbn string) (models.Book, error) {
	db := database.GetDbConnection()
	var book models.Book
	result := db.Model(models.Book{}).Find(&book, "isbn = ?", isbn)
	return book, result.Error
}

func (bs BookService) FindByTitle(title string) (models.Book, error) {
	db := database.GetDbConnection()
	var book models.Book
	result := db.Model(models.Book{}).Find(&book, "title = ?", title)
	return book, result.Error
}

func (bs BookService) FindAllByAuthor(author string) (models.Book, error) {
	db := database.GetDbConnection()
	var book models.Book
	result := db.Model(models.Book{}).Find(&book, "author = ?", author)
	return book, result.Error
}