package migrations

import (
	"ClothesShop/config"
	"ClothesShop/internal/models"
	"log"
)

func Migrate() {
	err := config.DB.AutoMigrate(
		&models.Product{},
		&models.Order{},
		&models.Cart{},
		&models.User{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate tables: %v", err)
	}
	log.Println("Database tables migrated successfully")
}
