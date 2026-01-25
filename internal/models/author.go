package models

import (
	"time"
)

type Author struct {
	ID			uint `gorm:"primaryKey"`
	Name		string
	Books		[]Book `gorm:"many2many:author_books"`
	CreatedAt	time.Time
	UpdatedAt	time.Time
}
