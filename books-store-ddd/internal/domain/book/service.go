package book

import "book-store/pkg/valueobjects"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateBook(book *Book) error {
	return s.repo.Save(book)
}

func (s *Service) GetBookByISBN(isbn valueobjects.ISBN) (*Book, error) {
	return s.repo.FindByISBN(isbn)
}
