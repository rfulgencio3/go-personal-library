package usecase

import (
	"github.com/rfulgencio3/go-personal-library/internal/domain"
	"github.com/rfulgencio3/go-personal-library/internal/repository"
	"github.com/rfulgencio3/go-personal-library/internal/validator"
)

type BookUseCase interface {
	CreateBook(book *domain.Book) error
	GetBookByID(id string) (*domain.Book, error)
	UpdateBook(book *domain.Book) error
	DeleteBook(id string) error
	GetAllBooks() ([]*domain.Book, error)
}

type bookUseCase struct {
	bookRepo repository.BookRepository
}

func NewBookUseCase(br repository.BookRepository) BookUseCase {
	return &bookUseCase{
		bookRepo: br,
	}
}

func (uc *bookUseCase) CreateBook(book *domain.Book) error {
	if err := validator.ValidateBook(book); err != nil {
		return err
	}
	return uc.bookRepo.Create(book)
}

func (uc *bookUseCase) GetBookByID(id string) (*domain.Book, error) {
	return uc.bookRepo.GetByID(id)
}

func (uc *bookUseCase) UpdateBook(book *domain.Book) error {
	if err := validator.ValidateBook(book); err != nil {
		return err
	}
	return uc.bookRepo.Update(book)
}

func (uc *bookUseCase) DeleteBook(id string) error {
	return uc.bookRepo.Delete(id)
}

func (uc *bookUseCase) GetAllBooks() ([]*domain.Book, error) {
	return uc.bookRepo.GetAll()
}
