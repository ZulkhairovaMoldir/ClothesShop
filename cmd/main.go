package main

import (
    "ClothesShop/config"
    "ClothesShop/internal/handlers"
    "ClothesShop/internal/repository"
    "ClothesShop/internal/services"
    "ClothesShop/middleware"
    "ClothesShop/migrations"
    "github.com/gin-gonic/gin"
    "log"
)

func main() {
    // Load environment variables and initialize DB
    config.LoadEnv()
    config.InitDB()

    // Run database migrations
    migrations.RunMigrations(config.DB)

    // Initialize User-related components
    userRepo := &repository.UserRepository{DB: config.DB}
    userService := &services.UserService{Repo: userRepo}
    userHandlers := &handlers.UserHandlers{Service: userService}
    authHandler := &handlers.AuthHandler{Service: userService}

    // Initialize Cart-related components
    cartRepo := &repository.CartRepository{DB: config.DB}
    cartService := &services.CartService{Repo: cartRepo}
    cartHandlers := &handlers.CartHandlers{Service: cartService}

    // Initialize Product-related components
    productRepo := &repository.ProductRepository{DB: config.DB}
    productService := &services.ProductService{Repo: productRepo}
    productHandlers := &handlers.ProductHandlers{Service: productService}

    // Initialize Order-related components
    orderRepo := &repository.OrderRepository{DB: config.DB}
    orderService := &services.OrderService{Repo: orderRepo}
    orderHandlers := &handlers.OrderHandlers{Service: orderService}

    router := gin.Default()
    router.Use(middleware.LoggingMiddleware())

    // Public routes
    public := router.Group("/")
    {
        public.POST("/register", userHandlers.CreateUser)
        public.POST("/login", authHandler.Login)

        // Public GET routes (accessible to unauthorized users)
        public.GET("/users", userHandlers.GetUsers)
        public.GET("/users/:id", userHandlers.GetUser)
        public.GET("/cart/:customerID", cartHandlers.GetCart)
        public.GET("/products", productHandlers.GetProducts)
        public.GET("/products/:id", productHandlers.GetProduct)
        public.GET("/orders", orderHandlers.GetOrders)
        public.GET("/orders/:id", orderHandlers.GetOrder)
    }

    // Protected routes (require authentication)
    protected := router.Group("/")
    protected.Use(middleware.AuthMiddleware())
    {
        // User routes
        protected.DELETE("/users/:id", userHandlers.DeleteUser)

        // Cart routes
        protected.POST("/cart", cartHandlers.AddItem)

        // Product routes
        protected.POST("/products", productHandlers.CreateProduct)
        protected.DELETE("/products/:id", productHandlers.DeleteProduct)

        // Order routes
        protected.POST("/orders", orderHandlers.CreateOrder)
    }

    // Health check route
    router.GET("/health-check", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "OK"})
    })

    // Start the server
    log.Println("Server running on port 8080...")
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}