
package repositories

import (
	"go_book_api/api/models"
	"gorm.io/gorm"
)

// Repository functions for Book
func CreateBook(db *gorm.DB, book *models.Book) error {
	return db.Create(book).Error
}

func GetBooks(db *gorm.DB) ([]models.Book, error) {
	var books []models.Book
	err := db.Find(&books).Error
	return books, err
}

func GetBook(db *gorm.DB, id interface{}) (*models.Book, error) {
	var book models.Book
	err := db.First(&book, "id = ?", id).Error
	return &book, err
}

func UpdateBook(db *gorm.DB, book *models.Book) error {
	return db.Save(book).Error
}

func DeleteBook(db *gorm.DB, id interface{}) error {
	var book models.Book
	return db.Delete(&book, id).Error
}