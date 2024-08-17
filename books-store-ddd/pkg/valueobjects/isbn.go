package valueobjects

import (
	"errors"
	"regexp"
)

type ISBN string

func NewISBN(value string) (ISBN, error) {
	if !isValidISBN(value) {
		return ISBN(""), errors.New("invalid ISBN format")
	}
	return ISBN(value), nil
}

func (i ISBN) String() string {
	return string(i)
}

func isValidISBN(isbn string) bool {
	isbnRegex := regexp.MustCompile(`^(ISBN(?:-10)?:? )?(?:\d{9}[\dX]|(?:ISBN(?:-13)?:? )?(?:978|979)[- ]?\d{1,5}[- ]?\d+[- ]?\d+[- ]?\d)$`)
	return isbnRegex.MatchString(isbn)
}
