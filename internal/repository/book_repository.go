package repository

import "github.com/rfulgencio3/go-personal-library/internal/domain"

type BookRepository interface {
	Create(book *domain.Book) error
	GetByID(id string) (*domain.Book, error)
	Update(book *domain.Book) error
	Delete(id string) error
	GetAll() ([]*domain.Book, error)
}
