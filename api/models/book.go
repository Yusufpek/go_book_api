package models

type Book struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
	PageCount int    `json:"page_count"`
	PublishedYear int    `json:"published_year"`
}

