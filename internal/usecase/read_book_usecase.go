package usecase

import (
	"github.com/rfulgencio3/go-personal-library/internal/domain"
	"github.com/rfulgencio3/go-personal-library/internal/repository"
)

type ReadBookUseCase interface {
	CreateReadBook(readBook *domain.ReadBook) error
	GetReadBookByID(id string) (*domain.ReadBook, error)
	GetAllReadBooks() ([]*domain.ReadBook, error)
	UpdateReadBook(readBook *domain.ReadBook) error
	DeleteReadBook(id string) error
	AddCommentToReadBook(id, comment string) error
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

func (u *readBookUseCase) GetAllReadBooks() ([]*domain.ReadBook, error) {
	return u.repo.GetAll()
}

func (u *readBookUseCase) UpdateReadBook(readBook *domain.ReadBook) error {
	return u.repo.Update(readBook)
}

func (u *readBookUseCase) DeleteReadBook(id string) error {
	return u.repo.Delete(id)
}

func (u *readBookUseCase) AddCommentToReadBook(id, comment string) error {
	return u.repo.AddComment(id, comment)
}
