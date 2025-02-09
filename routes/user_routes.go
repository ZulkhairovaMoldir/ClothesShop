package routes

import (
    "ClothesShop/internal/handlers"
    "ClothesShop/middleware"
    "github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.RouterGroup, userHandlers *handlers.UserHandlers, authHandler *handlers.AuthHandler) {
    userRoutes := router.Group("/users")
    {
        userRoutes.POST("", userHandlers.CreateUser) // Registration
        userRoutes.POST("/login", authHandler.Login) // Login
        userRoutes.GET("", middleware.AuthMiddleware(), userHandlers.GetUsers)
        userRoutes.GET("/:id", middleware.AuthMiddleware(), userHandlers.GetUser)
        userRoutes.DELETE("/:id", middleware.AuthMiddleware(), userHandlers.DeleteUser)
    }
}