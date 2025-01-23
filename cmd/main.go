package main

import (
	"ClothesShop/config"
	"ClothesShop/internal/handlers"
	"ClothesShop/internal/repository"
	"ClothesShop/internal/services"
	"ClothesShop/migrations"
	"ClothesShop/routes"
	"log"
)

func main() {
	config.InitDB()

	migrations.Migrate()

	userRepo := &repository.UserRepository{
		DB: config.DB,
	}

	userService := &services.UserService{
		Repo: userRepo,
	}

	userHandlers := &handlers.UserHandlers{
		Service: userService,
	}

	router := routes.SetupRoutes(userHandlers)

	log.Println("Сервер запущен на порту 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
