
package routes

import (
	"ClothesShop/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(router *gin.Engine, productHandlers *handlers.ProductHandlers) {
	productRoutes := router.Group("/products")
	{
		productRoutes.POST("", productHandlers.CreateProduct)
		productRoutes.GET("", productHandlers.GetProducts)
		productRoutes.GET("/:id", productHandlers.GetProduct)
		productRoutes.DELETE("/:id", productHandlers.DeleteProduct)
	}
}
