package handlers

import (
    "ClothesShop/internal/models"
    "ClothesShop/internal/services"
    "github.com/gin-contrib/sessions"
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
    "strconv"
)

type CartHandlers struct {
    Service *services.CartService
}

func (h *CartHandlers) GetCart(c *gin.Context) {
    session := sessions.Default(c)
    cart := session.Get("cart")
    if cart == nil {
        cart = []models.Cart{}
    }

    c.JSON(http.StatusOK, cart)
}

// AddItem adds an item to the cart
func (h *CartHandlers) AddItem(c *gin.Context) {
    var request struct {
        ProductID uint `json:"productId"`
    }
    if err := c.ShouldBindJSON(&request); err != nil {
        log.Printf("Error binding JSON: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    session := sessions.Default(c)
    cart := session.Get("cart")
    if cart == nil {
        cart = []models.Cart{}
    }

    cartItem := models.Cart{
        ProductID: request.ProductID,
        Quantity:  1, // Default quantity to 1
        CustomerID: 1, // Assuming a default customer ID for now
    }

    cart = append(cart.([]models.Cart), cartItem)
    session.Set("cart", cart)
    session.Save()

    c.JSON(http.StatusCreated, gin.H{"message": "Item added successfully"})
}

// RemoveItem removes an item from the cart
func (h *CartHandlers) RemoveItem(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        log.Printf("Error parsing cart ID: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Cart ID"})
        return
    }

    session := sessions.Default(c)
    cart := session.Get("cart")
    if cart == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
        return
    }

    cartItems := cart.([]models.Cart)
    for i, item := range cartItems {
        if item.ID == uint(id) {
            cartItems = append(cartItems[:i], cartItems[i+1:]...)
            break
        }
    }
    session.Set("cart", cartItems)
    session.Save()

    c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart"})
}