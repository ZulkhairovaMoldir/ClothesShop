package routes

import (
	"ClothesShop/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(router *gin.Engine, orderHandlers *handlers.OrderHandlers) {
	orderRoutes := router.Group("/orders")
	{
		orderRoutes.POST("", orderHandlers.CreateOrder)
		orderRoutes.GET("", orderHandlers.GetOrders)
		orderRoutes.GET("/:id", orderHandlers.GetOrder)
		orderRoutes.DELETE("/:id", orderHandlers.DeleteOrder)
	}
}
