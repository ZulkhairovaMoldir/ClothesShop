package models

import "gorm.io/gorm"

type Cart struct {
    ID         uint `gorm:"primaryKey"`
    ProductID  uint `gorm:"not null"`
    Quantity   int  `gorm:"not null"`
    CustomerID *uint `gorm:"default:null"` // Allow NULL values
    SessionID  *string `gorm:"default:null"` // Add SessionID for guest users
    CreatedAt  *gorm.DeletedAt
    UpdatedAt  *gorm.DeletedAt
}