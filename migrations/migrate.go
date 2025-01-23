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
	)
	if err != nil {
		log.Fatalf("Ошибка миграции таблиц: %v", err)
	}
	log.Println("Миграция таблиц выполнена успешно")
}
