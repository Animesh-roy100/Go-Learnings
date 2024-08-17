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
