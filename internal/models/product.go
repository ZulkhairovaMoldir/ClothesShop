package models

import "time"

type Product struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(100);not null"`
	Category    string `gorm:"type:varchar(50);not null"`
	Description string `gorm:"type:text"`
	Price       float64
	Stock       int
	ImageURL    string `json:"image_url"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Product) TableName() string {
	return "public.products"
}
