package routes

import (
	"ClothesShop/internal/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()

	// Define routes for products, orders, and carts
	router.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/orders", controllers.CreateOrder).Methods("POST")
	router.HandleFunc("/cart", controllers.GetCart).Methods("GET")

	// Return the router
	return router
}
