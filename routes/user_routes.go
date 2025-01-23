package routes

import (
	"ClothesShop/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	// Создаем группу маршрутов для пользователей
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", handlers.GetUsers)
		userRoutes.GET("/:id", handlers.GetUser)
		userRoutes.POST("/", handlers.CreateUser)
		userRoutes.DELETE("/:id", handlers.DeleteUser)
	}
}
