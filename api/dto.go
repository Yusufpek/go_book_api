package api

type CreateBookRequest struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	PageCount int    `json:"page_count" binding:"required"`
	PublishedYear int    `json:"published_year" default:"now()"`
}

type UpdateBookRequest struct {
	Title  *string `json:"title"`
	Author *string `json:"author"`
	PageCount *int    `json:"page_count"`
	PublishedYear *int    `json:"published_year"`
}