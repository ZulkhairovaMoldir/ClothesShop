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
    userID, exists := c.Get("userID")

    if exists {
        // LOGGED-IN USER → Fetch cart from the database
        cart, err := h.Service.GetCart(userID.(uint))
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch cart"})
            return
        }
        c.JSON(http.StatusOK, cart)
    } else {
        // GUEST USER → Fetch cart from session
        session := sessions.Default(c)
        cart := session.Get("cart")
        if cart == nil {
            cart = []models.Cart{}
        }
        c.JSON(http.StatusOK, cart)
    }
}

func (h *CartHandlers) AddItem(c *gin.Context) {
    var request struct {
        ProductID uint `json:"productId"`
    }
    if err := c.ShouldBindJSON(&request); err != nil {
        log.Printf("Error binding JSON: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    userID, exists := c.Get("userID")

    if exists {
        // LOGGED-IN USER → Save to Database
        userIDUint := userID.(uint)
        cartItem := models.Cart{
            CustomerID: &userIDUint, // Convert uint to *uint
            ProductID:  request.ProductID,
            Quantity:   1,
        }
        if err := h.Service.AddToCart(&cartItem); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not add to cart"})
            return
        }
    } else {
        // GUEST USER → Save to Session
        session := sessions.Default(c)
        cart := session.Get("cart")
        if cart == nil {
            cart = []models.Cart{}
        }
        cart = append(cart.([]models.Cart), models.Cart{ProductID: request.ProductID, Quantity: 1})
        session.Set("cart", cart)
        session.Save()
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Item added successfully"})
}

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