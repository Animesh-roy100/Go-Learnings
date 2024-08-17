package book

import (
	"book-store/pkg/valueobjects"
)

type Book struct {
	ID     valueobjects.ID
	Title  string
	Author string
	Price  valueobjects.Money
	ISBN   valueobjects.ISBN
}

func NewBook(title, author string, price valueobjects.Money, isbn valueobjects.ISBN) *Book {
	return &Book{
		ID:     valueobjects.NewID(),
		Title:  title,
		Author: author,
		Price:  price,
		ISBN:   isbn,
	}
}
