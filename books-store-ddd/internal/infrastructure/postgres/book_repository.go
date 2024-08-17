package postgres

import (
	"book-store/internal/domain/book"
	"book-store/pkg/valueobjects"
	"database/sql"

	"github.com/Masterminds/squirrel"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (cmd *BookRepository) Save(book *book.Book) error {
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	query := psql.Insert("books").Columns(
		"id", "title", "author", "price", "isbn").Values(
		book.ID, book.Title, book.Author, book.Price.Amount(), book.ISBN)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = cmd.db.Exec(sql, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *BookRepository) FindByID(id valueobjects.ID) (*book.Book, error) {
	// Implementation here
	return nil, nil
}

func (r *BookRepository) FindByISBN(isbn valueobjects.ISBN) (*book.Book, error) {
	// Implementation here
	return nil, nil
}

func (r *BookRepository) List(limit, offset int) ([]*book.Book, error) {
	// Implementation here
	return nil, nil
}

func (r *BookRepository) Update(book *book.Book) error {
	// Implementation here
	return nil
}

func (r *BookRepository) Delete(id valueobjects.ID) error {
	// Implementation here
	return nil
}
