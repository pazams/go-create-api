package models

// BookResponse is the book response model
type BookResponse struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// BooksResponse is the books response model
type BooksResponse []BookResponse
