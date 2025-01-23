package models

type User struct {
	ID       uint   `gorm:"primaryKey"`                        // Unique identifier
	Name     string `gorm:"type:varchar(100)"`                 // User's name
	Email    string `gorm:"type:varchar(100);unique;not null"` // Email, unique
	Password string `gorm:"type:varchar(100)"`                 // Password (hashed)
}
