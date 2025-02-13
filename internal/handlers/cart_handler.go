package handlers

import (
	"ClothesShop/internal/models"
	"ClothesShop/internal/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type CartHandlers struct {
	Service *services.CartService
}

func (h *CartHandlers) GetCart(c *gin.Context) {
	customerID, err := strconv.Atoi(c.Param("customerID"))
	if err != nil {
		log.Printf("Error parsing customerID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Customer ID"})
		return
	}
	log.Printf("Fetching cart for customerID: %d", customerID)

	cart, err := h.Service.GetCart(uint(customerID))
	if err != nil {
		log.Printf("Error fetching cart from service for customerID %d: %v", customerID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch cart"})
		return
	}
	log.Printf("Successfully retrieved cart for customerID %d: %v", customerID, cart)

	c.JSON(http.StatusOK, cart)
}

// AddItem adds an item to the cart
func (h *CartHandlers) AddItem(c *gin.Context) {
	var cartItem models.Cart
	if err := c.ShouldBindJSON(&cartItem); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	log.Printf("Adding item to cart: %v", cartItem)

	if err := h.Service.AddToCart(&cartItem); err != nil {
		log.Printf("Error adding item to cart: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to add item to cart"})
		return
	}
	log.Printf("Successfully added item to cart: %v", cartItem)

	c.JSON(http.StatusCreated, cartItem)
}

// RemoveItem removes an item from the cart
func (h *CartHandlers) RemoveItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Error parsing cart ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Cart ID"})
		return
	}
	log.Printf("Removing item with ID: %d", id)

	if err := h.Service.RemoveFromCart(uint(id)); err != nil {
		log.Printf("Error removing item with ID %d from cart: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to remove item from cart"})
		return
	}
	log.Printf("Successfully removed item with ID: %d", id)

	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart"})
}
