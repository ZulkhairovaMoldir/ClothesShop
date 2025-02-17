package routes

import (
    "ClothesShop/internal/handlers"
    "github.com/gin-gonic/gin"
)

func SetupCartRoutes(router *gin.Engine, cartHandlers *handlers.CartHandlers) {
    cartRoutes := router.Group("/cart")
    {
        cartRoutes.POST("/add", cartHandlers.AddItem)
        cartRoutes.GET("", cartHandlers.GetCart)
        cartRoutes.POST("/update", cartHandlers.UpdateItemQuantity) // Ensure this line is present
        cartRoutes.DELETE("/remove/:id", cartHandlers.RemoveItem)    // Ensure this line is present
    }
}