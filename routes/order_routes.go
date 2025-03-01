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
        orderRoutes.GET("", orderHandlers.GetOrders)    
        orderRoutes.GET("/:id", orderHandlers.GetOrder)
        orderRoutes.GET("/user", orderHandlers.GetOrdersByUser) // Add this line
	}
}
