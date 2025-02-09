package main

import (
    "ClothesShop/config"
    "log"

    "ClothesShop/internal/handlers"
    "ClothesShop/internal/repository"
    "ClothesShop/internal/services"
    "ClothesShop/middleware"
    "ClothesShop/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    // Load environment variables before anything else
    config.LoadEnv()

    config.InitDB()

    userRepo := &repository.UserRepository{
        DB: config.DB,
    }
    userService := &services.UserService{
        Repo: userRepo,
    }
    userHandlers := &handlers.UserHandlers{
        Service: userService,
    }

    authHandler := &handlers.AuthHandler{
        Service: userService,
    }

    router := gin.Default()
    router.Use(middleware.LoggingMiddleware())

    // Public routes
    public := router.Group("/")
    {
        public.POST("/register", userHandlers.CreateUser) // Registration
        public.POST("/login", authHandler.Login)          // Login
    }

    // Protected routes
    protected := router.Group("/")
    protected.Use(middleware.AuthMiddleware())
    {
        routes.SetupUserRoutes(protected, userHandlers, authHandler)
    }

    router.GET("/health-check", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "status": "OK",
        })
    })

    log.Println("Server running on port 8080...")
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}