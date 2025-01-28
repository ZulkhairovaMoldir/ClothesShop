package routes

import (
	"ClothesShop/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine, userHandlers *handlers.UserHandlers) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("", userHandlers.CreateUser)
		userRoutes.GET("", userHandlers.GetUsers)
		userRoutes.GET("/:id", userHandlers.GetUser)
		userRoutes.DELETE("/:id", userHandlers.DeleteUser)
	}
}
