package models

import "time"

type Cart struct {
    ID         uint       `gorm:"primaryKey"`
    ProductID  uint       `gorm:"not null"`
    Quantity   int        `gorm:"not null"`
    CustomerID *uint      `gorm:"default:null"` 
    SessionID  *string    `gorm:"default:null"`
    CreatedAt  time.Time
    UpdatedAt  time.Time
}