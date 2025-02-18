package models

import (
	"time"
)

type Order struct {
	ID          uint `gorm:"primaryKey"`
	CustomerID  uint `gorm:"not null"` // Ссылка на клиента
	TotalAmount float64
	Status      string     `gorm:"type:varchar(50);not null"`
	CreatedAt   time.Time  `gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime"`
	DeletedAt   *time.Time `gorm:"index"`
}
