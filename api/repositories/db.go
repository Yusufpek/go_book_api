package repositories

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	
	"log"
	"os"
	"go_book_api/api/models"
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

	err = DB.AutoMigrate(&models.Book{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}