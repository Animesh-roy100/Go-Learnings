package command

import (
	"book-store/internal/domain/book"
	"book-store/pkg/valueobjects"
)

type CreateBook struct {
	Title  string
	Author string
	Price  float64
	ISBN   string
}

type CreateBookHandler struct {
	bookService *book.Service
}

func NewCreateBookHandler(bookService *book.Service) *CreateBookHandler {
	return &CreateBookHandler{
		bookService: bookService,
	}
}

func (h *CreateBookHandler) Handle(cmd CreateBook) error {
	price, err := valueobjects.NewMoney(cmd.Price, "INR")
	if err != nil {
		return err
	}
	isbn, err := valueobjects.NewISBN(cmd.ISBN)
	if err != nil {
		return err
	}

	newBook := book.NewBook(cmd.Title, cmd.Author, price, isbn)
	return h.bookService.CreateBook(newBook)
}
