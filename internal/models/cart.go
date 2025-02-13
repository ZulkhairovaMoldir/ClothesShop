package models

import "gorm.io/gorm"

type Cart struct {
	ID         uint `gorm:"primaryKey"`
	ProductID  uint `gorm:"not null"` 
	Quantity   int  `gorm:"not null"`
	CustomerID uint `gorm:"not null"` 
	CreatedAt  *gorm.DeletedAt
	UpdatedAt  *gorm.DeletedAt
}