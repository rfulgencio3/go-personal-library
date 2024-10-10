package repository

import (
	"github.com/rfulgencio3/go-personal-library/internal/domain"
)

// ReadBookRepository defines the interface for operations on the ReadBook entity.
type ReadBookRepository interface {
	Create(readBook *domain.ReadBook) error
	GetByID(id string) (*domain.ReadBook, error)
	Update(readBook *domain.ReadBook) error
	Delete(id string) error
	GetAll() ([]*domain.ReadBook, error)
}
