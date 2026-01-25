package author

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

func (r *Repository) Create(author *models.Author) error {
	return r.db.Create(author).Error
}

func (r *Repository) GetByID(id uint) (*models.Author, error) {
	var author models.Author
	err := r.db.First(&author, id).Error
	return &author, err
}

func (r *Repository) GetAll() ([]models.Author, error) {
	var authors []models.Author
	err := r.db.Find(&authors).Error
	return authors, err
}

func (r *Repository) Update(author *models.Author) error {
	return r.db.Save(author).Error
}

func (r *Repository) Delete(id uint) error {
	return r.db.Delete(&models.Author{}, id).Error
}

// Additional method to get author with books preloaded
func (r *Repository) GetByIDWithBooks(id uint) (*models.Author, error) {
	var author models.Author
	err := r.db.Preload("Books").First(&author, id).Error
	return &author, err
}
