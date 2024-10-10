package validator

import (
	"errors"
	"strings"

	"github.com/rfulgencio3/go-personal-library/internal/domain"
)

var (
	ErrInvalidBookData = errors.New("invalid book data")
)

func ValidateBook(book *domain.Book) error {
	if strings.TrimSpace(book.Title) == "" {
		return errors.New("title is required")
	}
	if strings.TrimSpace(book.Author) == "" {
		return errors.New("author is required")
	}
	if book.Pages <= 0 {
		return errors.New("pages must be greater than zero")
	}

	return nil
}
