package api

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	
	"log"
	"os"
	"net/http"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	dsn := os.Getenv("DB_DSN")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) 
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.AutoMigrate(&Book{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}

func CreateBook(c *gin.Context) {
	var book Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		ResponseJson(c, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	DB.Create(&book)
	ResponseJson(c, http.StatusCreated, "Book created successfully", book)

}

func GetBooks(c *gin.Context) {
	var books []Book
	DB.Find(&books)
	ResponseJson(c, http.StatusOK, "Books retrieved successfully", books)
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	var book Book
	
	result := DB.First(&book, "id = ?", id)
	if result.Error != nil {
		ResponseJson(c, http.StatusNotFound, "Book not found", nil)
		return
	}
	
	ResponseJson(c, http.StatusOK, "Book retrieved successfully", book)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book Book
	
	result := DB.First(&book, "id = ?", id)
	if result.Error != nil {
		ResponseJson(c, http.StatusNotFound, "Book not found", nil)
		return
	}

	err := c.ShouldBindJSON(&book)
	if err != nil {
		ResponseJson(c, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	DB.Save(&book)
	ResponseJson(c, http.StatusOK, "Book updated successfully", book)
}

func DeleteBook(c *gin.Context) {
	var book Book
	id := c.Param("id")

	err := DB.Delete(&book, id).Error
	if err != nil {
		ResponseJson(c, http.StatusNotFound, "Book not found", nil)
		return
	}

	ResponseJson(c, http.StatusOK, "Book deleted successfully", nil)
}
