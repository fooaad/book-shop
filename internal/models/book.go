package models

import (
	"time"
)

type Book struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	Price     float64
	Stock     int
	Authors   []Author `gorm:"many2many:author_books;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
