package handlers

import (
    "ClothesShop/internal/models"
    "ClothesShop/internal/services"
    "crypto/rand"
    "encoding/hex"
    "github.com/gin-contrib/sessions"
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
    "strconv"
)

type CartHandlers struct {
    Service *services.CartService
}

// GetCart - Retrieves the cart for logged-in or guest users
func (h *CartHandlers) GetCart(c *gin.Context) {
    session := sessions.Default(c)
    userID, userExists := c.Get("userID")
    sessionID, sessionExists := session.Get("sessionID").(string)

    var cartItems []map[string]interface{}
    var err error

    if userExists {
        cartItems, err = h.Service.GetCart(userID.(uint))
    } else if sessionExists {
        cartItems, err = h.Service.GetGuestCart(sessionID)
    }

    if err != nil {
        log.Printf("Error fetching cart: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart"})
        return
    }

    c.JSON(http.StatusOK, cartItems) // Send only the array, not an object
}

// AddItem - Adds an item to the cart for logged-in or guest users
func (h *CartHandlers) AddItem(c *gin.Context) {
    var cartItem models.Cart

    if err := c.ShouldBindJSON(&cartItem); err != nil {
        log.Printf("Invalid request data: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    session := sessions.Default(c)
    userID, userExists := c.Get("userID")
    sessionID, sessionExists := session.Get("sessionID").(string)

    if userExists {
        cartItem.CustomerID = new(uint)
        *cartItem.CustomerID = userID.(uint)
    } else {
        if !sessionExists {
            sessionID = generateSessionID()
            session.Set("sessionID", sessionID)
            session.Save()
        }
        cartItem.SessionID = &sessionID
    }

    if err := h.Service.AddToCart(&cartItem); err != nil {
        log.Printf("Failed to add item to cart: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add to cart"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Item added successfully"})
}

// RemoveItem - Removes an item from the cart
func (h *CartHandlers) RemoveItem(c *gin.Context) {
    cartItemID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        log.Printf("Invalid cart ID format: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Cart ID"})
        return
    }

    session := sessions.Default(c)
    userID, userExists := c.Get("userID")

    if userExists {
        if err := h.Service.RemoveFromUserCart(uint(cartItemID), userID.(uint)); err != nil {
            log.Printf("Failed to remove item from user cart: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove item"})
            return
        }
    } else {
        cart := session.Get("cart")
        if cart == nil {
            log.Printf("Cart not found in session")
            c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
            return
        }

        cartItems, ok := cart.([]map[string]interface{})
        if !ok {
            log.Printf("Invalid cart format")
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid cart format"})
            return
        }

        updatedCart := make([]map[string]interface{}, 0)
        for _, item := range cartItems {
            if uint(item["id"].(float64)) != uint(cartItemID) {
                updatedCart = append(updatedCart, item)
            }
        }

        session.Set("cart", updatedCart)
        session.Save()
    }

    c.JSON(http.StatusOK, gin.H{"message": "Item removed successfully"})
}

// GenerateSessionID - Creates a secure session ID for guests
func generateSessionID() string {
    bytes := make([]byte, 16)
    _, err := rand.Read(bytes)
    if err != nil {
        log.Printf("Error generating session ID: %v", err)
        return ""
    }
    return hex.EncodeToString(bytes)
}