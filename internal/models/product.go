package models

import "gorm.io/gorm"

type Product struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(100);not null"`
	Category    string `gorm:"type:varchar(50);not null"`
	Description string `gorm:"type:text"`
	Price       float64
	Stock       int // Количество на складе
	CreatedAt   *gorm.DeletedAt
	UpdatedAt   *gorm.DeletedAt
}
