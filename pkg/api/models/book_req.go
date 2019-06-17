package models

// BookInsertRequest is the book insert request model
type BookInsertRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
