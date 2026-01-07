package main

import (
	"github.com/gin-gonic/gin"
	"go_book_api/api"
)

func main() {
	api.InitDB()

	r := gin.Default()

	r.POST("/book", api.CreateBook)
	r.GET("/books", api.GetBooks)
	r.GET("/books/:id", api.GetBook)
	r.PUT("/books/:id", api.UpdateBook)
	r.DELETE("/books/:id", api.DeleteBook)

	r.Run(":8000")
}