package routes

import (
	"ClothesShop/internal/handlers"
	"ClothesShop/middleware"
	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(router *gin.Engine, orderHandlers *handlers.OrderHandlers) {
	orderRoutes := router.Group("/orders")
	{
		orderRoutes.POST("", middleware.AuthMiddleware(), orderHandlers.CreateOrder) // Require login
		orderRoutes.GET("", middleware.AuthMiddleware(), orderHandlers.GetOrders)    // User can see their own orders
		orderRoutes.GET("/:id", middleware.AuthMiddleware(), orderHandlers.GetOrder) // User can see their order
	}
}
