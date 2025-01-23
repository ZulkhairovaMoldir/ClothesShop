package models

import "gorm.io/gorm"

type Order struct {
	ID          uint `gorm:"primaryKey"`
	CustomerID  uint `gorm:"not null"` // Ссылка на клиента
	TotalAmount float64
	Status      string `gorm:"type:varchar(50);not null"`
	CreatedAt   *gorm.DeletedAt
	UpdatedAt   *gorm.DeletedAt
}
