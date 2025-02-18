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

func (h *CartHandlers) AddItem(c *gin.Context) {
	var req struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	}

	// Читаем JSON и проверяем
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("❌ Ошибка разбора JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат запроса"})
		return
	}

	log.Println("🚀 Получен запрос на добавление в корзину:", req)

	// Проверяем переданный product_id
	if req.ProductID == 0 {
		log.Println("❌ Ошибка: передан неверный product_id:", req.ProductID)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный product_id"})
		return
	}

	session := sessions.Default(c)
	customerID, exists := c.Get("customerID")

	var cart models.Cart
	cart.ProductID = req.ProductID
	cart.Quantity = req.Quantity

	if exists {
		customerIDValue := customerID.(uint)
		cart.CustomerID = &customerIDValue
		cart.SessionID = nil
	} else {
		sessionID := session.Get("sessionID")
		if sessionID == nil {
			sessionID = h.GenerateSessionID()
			session.Set("sessionID", sessionID)
			session.Save()
		}
		sessionIDStr := sessionID.(string)
		cart.SessionID = &sessionIDStr
	}

	err := h.Service.AddToCart(&cart)
	if err != nil {
		log.Println("❌ Ошибка при добавлении в корзину:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("✅ Товар успешно добавлен в корзину")
	c.JSON(http.StatusOK, gin.H{"message": "Товар добавлен в корзину"})
}

func (h *CartHandlers) UpdateItemQuantity(c *gin.Context) {
	var req struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := sessions.Default(c)
	customerID, exists := c.Get("customerID")

	var sessionID *string
	var customerIDPtr *uint

	if exists {
		id := customerID.(uint)
		customerIDPtr = &id
	} else {
		sessionIDVal := session.Get("sessionID")
		if sessionIDVal != nil {
			sessionIDStr := sessionIDVal.(string)
			sessionID = &sessionIDStr
		}
	}

	updatedCartItem, err := h.Service.UpdateCartQuantity(req.ProductID, req.Quantity, sessionID, customerIDPtr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if updatedCartItem == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Quantity updated", "quantity": updatedCartItem.Quantity})
}

func (h *CartHandlers) RemoveItem(c *gin.Context) {
	productID, _ := strconv.Atoi(c.Param("id"))
	session := sessions.Default(c)
	customerID, exists := c.Get("customerID")

	if exists {
		if err := h.Service.RemoveFromUserCart(uint(productID), customerID.(uint)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		sessionID := session.Get("sessionID")
		if sessionID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No session found"})
			return
		}
		if err := h.Service.RemoveFromGuestCart(uint(productID), sessionID.(string)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item removed"})
}

func (h *CartHandlers) GetCart(c *gin.Context) {
	session := sessions.Default(c)
	customerID, exists := c.Get("customerID")

	var cartItems []map[string]interface{}
	var err error

	if exists {
		cartItems, err = h.Service.GetCart(customerID.(uint))
	} else {
		sessionID := session.Get("sessionID")
		if sessionID == nil {
			sessionID = h.GenerateSessionID()
			session.Set("sessionID", sessionID)
			session.Save()
		}
		cartItems, err = h.Service.GetGuestCart(sessionID.(string))
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cartItems)
}

func (h *CartHandlers) GenerateSessionID() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(b)
}
