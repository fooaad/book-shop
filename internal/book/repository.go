package book

import (
	"book-shop/internal/models"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(book *models.Book) error {
	return r.db.Create(book).Error
}

func (r *Repository) GetByID(id uint) (*models.Book, error) {
	var book models.Book
	err := r.db.First(&book, id).Error
	return &book, err
}

func (r *Repository) GetAll() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *Repository) Update(book *models.Book) error {
	return r.db.Save(book).Error
}

func (r *Repository) Delete(id uint) error {
	return r.db.Delete(&models.Book{}, id).Error
}

// Additional method to get book with authors preloaded
func (r *Repository) GetByIDWithAuthors(id uint) (*models.Book, error) {
	var book models.Book
	err := r.db.Preload("Authors").First(&book, id).Error
	return &book, err
}
