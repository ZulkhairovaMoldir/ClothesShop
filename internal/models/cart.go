package models

import "time"

type Cart struct {
    ID         uint   `json:"id" gorm:"primaryKey"`
    ProductID  uint   `json:"product_id"`
    Quantity   int    `json:"quantity"`
    CustomerID *uint  `json:"customer_id,omitempty"`
    SessionID  *string `json:"session_id,omitempty"`
    CreatedAt  time.Time
    UpdatedAt  time.Time
}