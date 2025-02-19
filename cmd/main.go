package main

import (
    "ClothesShop/config"
    "ClothesShop/internal/handlers"
    "ClothesShop/internal/repository"
    "ClothesShop/internal/services"
    "ClothesShop/middleware"
    "ClothesShop/migrations"
    "github.com/gin-contrib/cors"
    "github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/cookie"
    "github.com/gin-gonic/gin"
    "log"
    "time"
)

func main() {
    config.LoadEnv()
    config.InitDB()

    migrations.RunMigrations(config.DB)

    userRepo := &repository.UserRepository{DB: config.DB}
    cartRepo := &repository.CartRepository{DB: config.DB}
    cartService := &services.CartService{Repo: cartRepo}
    userService := &services.UserService{Repo: userRepo, CartService: cartService}
    userHandlers := &handlers.UserHandlers{Service: userService, CartService: cartService}
    authHandler := &handlers.AuthHandler{Service: userService, CartService: cartService}

    cartHandlers := &handlers.CartHandlers{Service: cartService}

    productRepo := &repository.ProductRepository{DB: config.DB}
    productService := &services.ProductService{Repo: productRepo}
    productHandlers := &handlers.ProductHandlers{Service: productService}

    orderRepo := &repository.OrderRepository{DB: config.DB}
    orderService := &services.OrderService{Repo: orderRepo}
    orderHandlers := &handlers.OrderHandlers{Service: orderService}

    router := gin.Default()

    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:8080", "http://localhost.:8080"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

    store := cookie.NewStore([]byte("secret"))
    router.Use(sessions.Sessions("mysession", store))

    router.Use(middleware.LoggingMiddleware())

    router.Static("/static", "./static")

    router.GET("/", func(c *gin.Context) {
        c.File("./static/index.html")
    })

    public := router.Group("/")
    {
        public.POST("/register", userHandlers.CreateUser)
        public.POST("/login", authHandler.Login)

        public.GET("/users", userHandlers.GetUsers)
        public.GET("/users/:id", userHandlers.GetUser)
        public.GET("/cart", cartHandlers.GetCart)
        public.POST("/cart/add", cartHandlers.AddItem)
        public.POST("/cart/update", cartHandlers.UpdateItemQuantity)
        public.DELETE("/cart/remove/:id", cartHandlers.RemoveItem)
        public.GET("/products", productHandlers.GetProducts)
        public.GET("/products/:id", productHandlers.GetProduct)
        public.GET("/orders", orderHandlers.GetOrders)
        public.GET("/orders/:id", orderHandlers.GetOrder)
    }

    protected := router.Group("/")
    protected.Use(middleware.AuthMiddleware())
    {
        protected.GET("/profile", userHandlers.GetProfile)
        protected.DELETE("/users/:id", userHandlers.DeleteUser)

        protected.DELETE("/cart/item/:id", cartHandlers.RemoveItem)

        protected.POST("/products", productHandlers.CreateProduct)
        protected.DELETE("/products/:id", productHandlers.DeleteProduct)

        protected.POST("/orders", orderHandlers.CreateOrder)
        protected.GET("/orders/user", orderHandlers.GetOrdersByUser) // Add this line
    }

    router.GET("/health-check", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "OK"})
    })

    log.Println("Server running on port 8080...")
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}