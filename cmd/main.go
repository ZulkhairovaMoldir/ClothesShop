package main

import (
	"ClothesShop/config"
	"ClothesShop/migrations"
	"ClothesShop/routes"
	"log"
)

func main() {
	config.InitDB()

	migrations.Migrate()

	router := routes.SetupRoutes()

	log.Println("Сервер запущен на порту 8080...")
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
