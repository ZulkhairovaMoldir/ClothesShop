package main

import (
	"ClothesShop/config"
	"ClothesShop/migrations"
	"ClothesShop/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	migrations.Migrate()

	router := gin.Default()

	routes.RegisterUserRoutes(router)

	router.Run(":8080")
}
