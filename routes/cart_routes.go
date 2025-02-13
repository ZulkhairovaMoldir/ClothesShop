package routes

import (
    "ClothesShop/internal/handlers"
    "github.com/gin-gonic/gin"
)

func SetupCartRoutes(router *gin.Engine, cartHandlers *handlers.CartHandlers) {
    cartRoutes := router.Group("/cart")
    {
        cartRoutes.POST("", cartHandlers.AddItem) // Public
        cartRoutes.GET("/:customerID", cartHandlers.GetCart) // Public (temporary)
        cartRoutes.DELETE("/item/:id", cartHandlers.RemoveItem) // Public (temporary)
    }
}