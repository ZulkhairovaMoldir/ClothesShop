package main

import (
    "ClothesShop/config"
    "ClothesShop/internal/handlers"
    "ClothesShop/internal/repository"
    "ClothesShop/internal/services"
    "ClothesShop/middleware"
    "ClothesShop/migrations"
    "ClothesShop/routes"
    "log"
    "github.com/gin-gonic/gin"
)

func main() {
    config.InitDB()

    migrations.Migrate()

    userRepo := &repository.UserRepository{
        DB: config.DB,
    }
    userService := &services.UserService{
        Repo: userRepo,
    }
    userHandlers := &handlers.UserHandlers{
        Service: userService,
    }

    orderRepo := &repository.OrderRepository{
        DB: config.DB,
    }
    orderService := &services.OrderService{
        Repo: orderRepo,
    }
    orderHandlers := &handlers.OrderHandlers{
        Service: orderService,
    }

    cartRepo := &repository.CartRepository{
        DB: config.DB,
    }
    cartService := &services.CartService{
        Repo: cartRepo,
    }
    cartHandlers := &handlers.CartHandlers{
        Service: cartService,
    }


    productRepo := &repository.ProductRepository{
        DB: config.DB,
    }
    productService := &services.ProductService{
        Repo: productRepo,
    }
    productHandlers := &handlers.ProductHandlers{
        Service: productService,
    }

 
    router := gin.Default()
    router.Use(middleware.LoggingMiddleware()) 
    router.Use(middleware.AuthMiddleware())    

    routes.SetupUserRoutes(router, userHandlers)        
    routes.SetupOrderRoutes(router, orderHandlers)       
    routes.SetupCartRoutes(router, cartHandlers)         
    routes.SetupProductRoutes(router, productHandlers)   

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