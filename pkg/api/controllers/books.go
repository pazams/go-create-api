package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/AndrewBurian/powermux"

	"github.com/pazams/go-create-api/pkg/api/data"
	"github.com/pazams/go-create-api/pkg/api/models"
)

// BookController ..
type BookController struct {
	dal *data.DAL
}

// NewBookController ..
func NewBookController(dal *data.DAL) *BookController {
	return &BookController{dal: dal}
}

// Book returns a book by id
func (c *BookController) Book(w http.ResponseWriter, r *http.Request) (int, interface{}) {
	id := powermux.PathParam(r, "id")
	book, err := c.dal.SelectBookByID(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, toBookResponse(book)
}

// Books ..
func (c *BookController) Books(w http.ResponseWriter, r *http.Request) (int, interface{}) {
	books, err := c.dal.SelectBooks()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, toBooksResponse(books)
}

// InsertBook ..
func (c *BookController) InsertBook(w http.ResponseWriter, r *http.Request) (int, interface{}) {
	// todo parse b
	reqModel := &models.BookInsertRequest{}
	err := json.NewDecoder(r.Body).Decode(reqModel)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	book, err := c.dal.InsertBook(toBook(reqModel))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, toBookResponse(book)
}

func toBook(b *models.BookInsertRequest) *data.Book {
	return &data.Book{
		Title:  b.Title,
		Author: b.Author,
	}
}

func toBookResponse(b *data.Book) *models.BookResponse {
	return &models.BookResponse{
		ID:     b.Id,
		Title:  b.Title,
		Author: b.Author,
	}
}

func toBooksResponse(bs []data.Book) models.BooksResponse {

	booksResponse := make(models.BooksResponse, 0, len(bs))

	for _, b := range bs {
		booksResponse = append(booksResponse, *toBookResponse(&b))
	}

	return booksResponse

}
