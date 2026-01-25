package author

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

func (s *Service) CreateAuthor(author *models.Author) error {
	if author.Name == "" {
		return errors.New("author name cannot be empty")
	}
	return s.repo.Create(author)
}

func (s *Service) GetAuthorByID(id uint) (*models.Author, error) {
	if id == 0 {
		return nil, errors.New("invalid author ID")
	}
	return s.repo.GetByID(id)
}

func (s *Service) GetAllAuthors() ([]models.Author, error) {
	return s.repo.GetAll()
}

func (s *Service) UpdateAuthor(author *models.Author) error {
	if author.ID == 0 {
		return errors.New("author ID is required for update")
	}
	if author.Name == "" {
		return errors.New("author name cannot be empty")
	}
	return s.repo.Update(author)
}

func (s *Service) DeleteAuthor(id uint) error {
	if id == 0 {
		return errors.New("invalid author ID")
	}
	return s.repo.Delete(id)
}

func (s *Service) GetAuthorByIDWithBooks(id uint) (*models.Author, error) {
	if id == 0 {
		return nil, errors.New("invalid author ID")
	}
	return s.repo.GetByIDWithBooks(id)
}
