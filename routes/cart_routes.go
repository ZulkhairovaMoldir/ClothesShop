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
        cartRoutes.POST("/update", cartHandlers.UpdateItemQuantity)
        cartRoutes.DELETE("/remove/:id", cartHandlers.RemoveItem)
    }
}