package handlers

import (
    "ClothesShop/internal/models"
    "ClothesShop/internal/services"
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