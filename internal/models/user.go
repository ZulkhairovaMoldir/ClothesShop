package models

import "gorm.io/gorm"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(100)"`
	Email     string `gorm:"type:varchar(100);unique;not null"`
	Password  string `gorm:"type:varchar(100)"`
	CreatedAt *gorm.DeletedAt
	UpdatedAt *gorm.DeletedAt
}
