package query

import (
	"book-store/internal/domain/book"
	"book-store/pkg/valueobjects"
)

type GetBookByISBN struct {
	ISBN string
}

type GetBookByISBNHandler struct {
	bookService *book.Service
}

func NewGetBookByISBNHandler(bookService *book.Service) *GetBookByISBNHandler {
	return &GetBookByISBNHandler{bookService: bookService}
}

func (h *GetBookByISBNHandler) Handle(query GetBookByISBN) (*book.Book, error) {
	isbn, err := valueobjects.NewISBN(query.ISBN)
	if err != nil {
		return nil, err
	}
	return h.bookService.GetBookByISBN(isbn)
}
