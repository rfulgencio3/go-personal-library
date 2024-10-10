package usecase

import (
	"github.com/rfulgencio3/go-personal-library/internal/domain"
	"github.com/rfulgencio3/go-personal-library/internal/repository"
)

// ReadBookUseCase defines the business logic for ReadBook operations.
type ReadBookUseCase interface {
	CreateReadBook(readBook *domain.ReadBook) error
	GetReadBookByID(id string) (*domain.ReadBook, error)
}

type readBookUseCase struct {
	repo repository.ReadBookRepository
}

func NewReadBookUseCase(repo repository.ReadBookRepository) ReadBookUseCase {
	return &readBookUseCase{repo: repo}
}

func (u *readBookUseCase) CreateReadBook(readBook *domain.ReadBook) error {
	return u.repo.Create(readBook)
}

func (u *readBookUseCase) GetReadBookByID(id string) (*domain.ReadBook, error) {
	return u.repo.GetByID(id)
}
