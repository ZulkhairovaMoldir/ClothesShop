package models

import "time"

type Order struct {
    ID          uint       `gorm:"primaryKey"`
    CustomerID  *uint      `gorm:"default:null"`
    SessionID   *string    `gorm:"default:null"`
    TotalAmount float64    `gorm:"not null"`
    Status      string     `gorm:"type:varchar(100)"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
}