package routes

import (
	"ClothesShop/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(userHandlers *handlers.UserHandlers) *gin.Engine {
	router := gin.Default()

	router.POST("/users", userHandlers.CreateUser)
	router.GET("/users", userHandlers.GetUsers)
	router.GET("/users/:id", userHandlers.GetUser)
	router.DELETE("/users/:id", userHandlers.DeleteUser)

	return router
}
