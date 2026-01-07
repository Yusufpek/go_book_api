
package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "go_book_api/api/models"
    "go_book_api/api/repositories"
	"go_book_api/api/response"
)

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		response.ResponseJson(c, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}
	if err := repositories.CreateBook(repositories.DB, &book); err != nil {
		response.ResponseJson(c, http.StatusInternalServerError, "Failed to create book", nil)
		return
	}
	response.ResponseJson(c, http.StatusCreated, "Book created successfully", book)
}

func GetBooks(c *gin.Context) {
	books, err := repositories.GetBooks(repositories.DB)
	if err != nil {
		response.ResponseJson(c, http.StatusInternalServerError, "Failed to retrieve books", nil)
		return
	}
	response.ResponseJson(c, http.StatusOK, "Books retrieved successfully", books)
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	book, err := repositories.GetBook(repositories.DB, id)
	if err != nil {
		response.ResponseJson(c, http.StatusNotFound, "Book not found", nil)
		return
	}
	response.ResponseJson(c, http.StatusOK, "Book retrieved successfully", book)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	book, err := repositories.GetBook(repositories.DB, id)
	if err != nil {
		response.ResponseJson(c, http.StatusNotFound, "Book not found", nil)
		return
	}
	if err := c.ShouldBindJSON(book); err != nil {
		response.ResponseJson(c, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}
	if err := repositories.UpdateBook(repositories.DB, book); err != nil {
		response.ResponseJson(c, http.StatusInternalServerError, "Failed to update book", nil)
		return
	}
	response.ResponseJson(c, http.StatusOK, "Book updated successfully", book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	if err := repositories.DeleteBook(repositories.DB, id); err != nil {
		response.ResponseJson(c, http.StatusNotFound, "Book not found", nil)
		return
	}
	response.ResponseJson(c, http.StatusOK, "Book deleted successfully", nil)
}
