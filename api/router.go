package api

import (
    "github.com/gin-gonic/gin"
    "go_book_api/api/handlers"
)

// NewRouter sets up all routes and returns a Gin engine
func NewRouter() *gin.Engine {
    r := gin.Default()
    r.POST("/book", handlers.CreateBook)
    r.GET("/books", handlers.GetBooks)
    r.GET("/books/:id", handlers.GetBook)
    r.PUT("/books/:id", handlers.UpdateBook)
    r.DELETE("/books/:id", handlers.DeleteBook)
    return r
}
