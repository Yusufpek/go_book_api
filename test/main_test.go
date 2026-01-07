package tests

import (
    "bytes"
    "encoding/json"
    "log"
    "net/http"
    "net/http/httptest"
    "strconv"
    "testing"

    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
	"gorm.io/gorm/logger"
    "gorm.io/gorm"

    "go_book_api/api"
)

func setupTestDB() {
	var err error
	api.DB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to test database: %v", err)
	}
	if err := api.DB.AutoMigrate(&api.Book{}); err != nil {
		log.Fatalf("Failed to migrate test database: %v", err)
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/book", api.CreateBook)
	r.GET("/books", api.GetBooks)
	r.GET("/books/:id", api.GetBook)
	r.PUT("/books/:id", api.UpdateBook)
	r.DELETE("/books/:id", api.DeleteBook)
	return r
}

func decodeResponse(w *httptest.ResponseRecorder, t *testing.T) api.JsonResponse {
	var response api.JsonResponse
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}
	return response
}

func addBook() api.Book {
	book := api.Book{
		Title: "Test Book",
		Author: "Test Author",
		PageCount: 123,
		PublishedYear: 2020,
	}

	api.DB.Create(&book)
	return book
}

func TestCreateBook(t *testing.T) {
    setupTestDB()
    router := setupRouter()

    book := api.Book{
        Title:         "New Book",
        Author:        "New Author",
        PageCount:     200,
        PublishedYear: 2021,
    }

    jsonValue, _ := json.Marshal(book)
    req, _ := http.NewRequest("POST", "/book", bytes.NewBuffer(jsonValue))
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    if w.Code != http.StatusCreated {
        t.Errorf("Expected status %d but got %d", http.StatusCreated, w.Code)
    }

    response := decodeResponse(w, t)
    if response.Data == nil {
        t.Errorf("Expected book data in response but got nil")
    }
}

func TestGetBooks(t *testing.T) {
    setupTestDB()
    addBook()
    router := setupRouter()

    req, _ := http.NewRequest("GET", "/books", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    if status := w.Code; status != http.StatusOK {
        t.Errorf("Expected status %d, got %d", http.StatusOK, status)
    }

    response := decodeResponse(w, t)
    books, ok := response.Data.([]interface{})
    if !ok || len(books) == 0 {
        t.Errorf("Expected non-empty books list")
    }
}

func TestGetBook(t *testing.T) {
    setupTestDB()
    book := addBook()
    router := setupRouter()

    req, _ := http.NewRequest("GET", "/books/"+strconv.Itoa(int(book.ID)), nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    if status := w.Code; status != http.StatusOK {
        t.Errorf("Expected status %d, got %d", http.StatusOK, status)
    }

    response := decodeResponse(w, t)
    data, ok := response.Data.(map[string]interface{})
    if !ok || int(data["id"].(float64)) != int(book.ID) {
        t.Errorf("Expected book ID %d, got %v", book.ID, data["id"])
    }
}

func TestUpdateBook(t *testing.T) {
    setupTestDB()
    book := addBook()
    router := setupRouter()

    updateBook := api.Book{
        Title:         "Advanced Go Programming",
        Author:        "Demo Author name",
        PageCount:     350,
        PublishedYear: 2021,
    }
    jsonValue, _ := json.Marshal(updateBook)

    req, _ := http.NewRequest("PUT", "/books/"+strconv.Itoa(int(book.ID)), bytes.NewBuffer(jsonValue))
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    if status := w.Code; status != http.StatusOK {
        t.Errorf("Expected status %d, got %d", http.StatusOK, status)
    }

    response := decodeResponse(w, t)
    data, ok := response.Data.(map[string]interface{})
    if !ok || data["title"] != "Advanced Go Programming" {
        t.Errorf("Expected updated book title 'Advanced Go Programming', got %v", response.Data)
    }
}

func TestDeleteBook(t *testing.T) {
    setupTestDB()
    book := addBook()
    router := setupRouter()
	
    req, _ := http.NewRequest("DELETE", "/books/"+strconv.Itoa(int(book.ID)), nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    if status := w.Code; status != http.StatusOK {
        t.Errorf("Expected status %d, got %d", http.StatusOK, status)
    }

    response := decodeResponse(w, t)
    if response.Message != "Book deleted successfully" {
        t.Errorf("Expected delete message 'Book deleted successfully', got %v", response.Message)
    }

    // Verify that the book was deleted
	api.DB.Logger = logger.Default.LogMode(logger.Silent)
    var deletedBook api.Book
    result := api.DB.First(&deletedBook, book.ID)
    if result.Error == nil {
        t.Errorf("Expected book to be deleted, but it still exists")
    }
}

