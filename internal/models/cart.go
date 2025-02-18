package models

import "time"

type Cart struct {
	ID         uint    `json:"id" gorm:"primaryKey"`
	ProductID  uint    `json:"product_id"`
	Product    Product `json:"product" gorm:"foreignKey:ProductID"`
	Quantity   int     `json:"quantity"`
	CustomerID *uint   `json:"customer_id,omitempty"`
	SessionID  *string `json:"session_id,omitempty"`
	Size       string  `json:"size"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
