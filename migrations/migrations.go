package main

import (
    "ClothesShop/config"
    "ClothesShop/internal/models"
    "log"
)

func main() {
    // Инициализация базы данных
    config.InitDB()

    // Выполнение миграций
    err := config.DB.AutoMigrate(
        &models.Product{},
        &models.Order{},
        &models.Cart{},
        &models.User{},
    )
    if err != nil {
        log.Fatalf("Failed to migrate tables: %v", err)
    }

    log.Println("Migration completed successfully")
}