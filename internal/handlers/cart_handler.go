package handlers

import (
	"ClothesShop/internal/models"
	"ClothesShop/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CartHandlers struct {
	Service *services.CartService
}

func (h *CartHandlers) GetCart(c *gin.Context) {
	customerID, err := strconv.Atoi(c.Param("customerID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Customer ID"})
		return
	}

	cart, err := h.Service.GetCart(uint(customerID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch cart"})
		return
	}
	c.JSON(http.StatusOK, cart)
}

func (h *CartHandlers) AddItem(c *gin.Context) {
	var cartItem models.Cart
	if err := c.ShouldBindJSON(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.Service.AddToCart(&cartItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to add item to cart"})
		return
	}
	c.JSON(http.StatusCreated, cartItem)
}

func (h *CartHandlers) RemoveItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Cart ID"})
		return
	}

	if err := h.Service.RemoveFromCart(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to remove item from cart"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart"})
}
