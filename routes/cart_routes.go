
package routes

import (
	"ClothesShop/internal/handlers"
	"ClothesShop/middleware"
	"github.com/gin-gonic/gin"
)

func SetupCartRoutes(router *gin.Engine, cartHandlers *handlers.CartHandlers) {
	cartRoutes := router.Group("/carts", middleware.AuthMiddleware())
	{
		cartRoutes.POST("", cartHandlers.AddItem)
		cartRoutes.GET("/:customerID", cartHandlers.GetCart)
		cartRoutes.DELETE("/:id", cartHandlers.RemoveItem)
	}
}
