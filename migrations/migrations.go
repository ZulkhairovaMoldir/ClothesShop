package migrations

import (
    "ClothesShop/internal/models"
    "gorm.io/gorm"
    "log"
)

func RunMigrations(db *gorm.DB) {
    log.Println("Running migrations...")
    err := db.AutoMigrate(&models.Product{}, &models.User{}, &models.Order{}, &models.Cart{})
    if err != nil {
        log.Fatalf("Migration error: %v", err)
    }
    log.Println("Migrations completed successfully!")
}