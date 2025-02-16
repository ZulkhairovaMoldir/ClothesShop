package routes

import (
	"ClothesShop/internal/handlers"
	"ClothesShop/middleware"
	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(router *gin.Engine, productHandlers *handlers.ProductHandlers) {
	productRoutes := router.Group("/products")
	{
		productRoutes.POST("", middleware.AuthMiddleware(), productHandlers.CreateProduct) // Only admins can add products
		productRoutes.GET("", productHandlers.GetProducts)                                 // Public
		productRoutes.GET("/:id", productHandlers.GetProduct)                              // Public
	}
}
