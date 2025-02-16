package routes

import (
	"ClothesShop/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupCartRoutes(router *gin.Engine, cartHandlers *handlers.CartHandlers) {
	cartRoutes := router.Group("/cart")
	{
		cartRoutes.POST("/add", cartHandlers.AddItem)             // Fix endpoint
		cartRoutes.GET("", cartHandlers.GetCart)                  // Fix for guests and users
		cartRoutes.DELETE("/remove/:id", cartHandlers.RemoveItem) // Fix endpoint
	}
}
