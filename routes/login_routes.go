package routes

import (
	"ClothesShop/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.Engine, authHandler *handlers.AuthHandler) {
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/login", authHandler.Login)
	}
}
