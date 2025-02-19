package handlers

import (
    "ClothesShop/internal/models"
    "ClothesShop/internal/services"
    "github.com/gin-contrib/sessions"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type OrderHandlers struct {
    Service *services.OrderService
}

func (h *OrderHandlers) GetOrders(c *gin.Context) {
    orders, err := h.Service.GetOrders()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch orders"})
        return
    }
    c.JSON(http.StatusOK, orders)
}

func (h *OrderHandlers) GetOrder(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    order, err := h.Service.GetOrder(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }
    c.JSON(http.StatusOK, order)
}

func (h *OrderHandlers) CreateOrder(c *gin.Context) {
    var order models.Order
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    session := sessions.Default(c)
    customerID, exists := c.Get("customerID")

    if exists {
        customerIDValue := customerID.(uint)
        order.CustomerID = &customerIDValue
        order.SessionID = nil
    } else {
        sessionID := session.Get("sessionID")
        if sessionID == nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "No session found"})
            return
        }
        sessionIDStr := sessionID.(string)
        order.SessionID = &sessionIDStr
    }

    if err := h.Service.CreateOrder(&order); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create order"})
        return
    }
    c.JSON(http.StatusCreated, order)
}

func (h *OrderHandlers) DeleteOrder(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := h.Service.DeleteOrder(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete order"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Order deleted"})
}

func (h *OrderHandlers) GetOrdersByUser(c *gin.Context) {
    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    orders, err := h.Service.GetOrdersByCustomerID(userID.(uint))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch orders"})
        return
    }
    c.JSON(http.StatusOK, orders)
}