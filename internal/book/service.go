package book

import (
	"book-shop/internal/models"
	"errors"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateBook(book *models.Book) error {
	if book.Title == "" {
		return errors.New("book title cannot be empty")
	}
	if book.Price < 0 {
		return errors.New("book price cannot be negative")
	}
	if book.Stock < 0 {
		return errors.New("book stock cannot be negative")
	}
	return s.repo.Create(book)
}

func (s *Service) GetBookByID(id uint) (*models.Book, error) {
	if id == 0 {
		return nil, errors.New("invalid book ID")
	}
	return s.repo.GetByID(id)
}

func (s *Service) GetAllBooks() ([]models.Book, error) {
	return s.repo.GetAll()
}

func (s *Service) UpdateBook(book *models.Book) error {
	if book.ID == 0 {
		return errors.New("book ID is required for update")
	}
	if book.Title == "" {
		return errors.New("book title cannot be empty")
	}
	if book.Price < 0 {
		return errors.New("book price cannot be negative")
	}
	if book.Stock < 0 {
		return errors.New("book stock cannot be negative")
	}
	return s.repo.Update(book)
}

func (s *Service) DeleteBook(id uint) error {
	if id == 0 {
		return errors.New("invalid book ID")
	}
	return s.repo.Delete(id)
}

func (s *Service) GetBookByIDWithAuthors(id uint) (*models.Book, error) {
	if id == 0 {
		return nil, errors.New("invalid book ID")
	}
	return s.repo.GetByIDWithAuthors(id)
}
