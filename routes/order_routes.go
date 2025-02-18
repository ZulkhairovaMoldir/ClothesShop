package routes

import (
	"ClothesShop/internal/handlers"
	"ClothesShop/middleware"
	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(router *gin.Engine, orderHandlers *handlers.OrderHandlers) {
	orderRoutes := router.Group("/orders")
	{
		orderRoutes.POST("", middleware.AuthMiddleware(), orderHandlers.CreateOrder)
		orderRoutes.GET("", middleware.AuthMiddleware(), orderHandlers.GetOrders)    
		orderRoutes.GET("/:id", middleware.AuthMiddleware(), orderHandlers.GetOrder) 
	}
}
