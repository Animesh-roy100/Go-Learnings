package book

import "book-store/pkg/valueobjects"

type Repository interface {
	Save(book *Book) error
	FindByID(id valueobjects.ID) (*Book, error)
	FindByISBN(isbn valueobjects.ISBN) (*Book, error)
	List(limit, offset int) ([]*Book, error)
	Update(book *Book) error
	Delete(id valueobjects.ID) error
}
