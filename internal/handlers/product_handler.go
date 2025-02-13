package handlers

import (
    "ClothesShop/internal/models"
    "ClothesShop/internal/services"
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
    "strconv"
)

type ProductHandlers struct {
    Service *services.ProductService
}

func (h *ProductHandlers) GetProducts(c *gin.Context) {
    log.Println("Handler: Fetching products...") // Debug log
    products, err := h.Service.GetProducts()
    if err != nil {
        log.Println("Error fetching products:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
        return
    }
    log.Println("Fetched products:", products) // Debugging log
    c.JSON(http.StatusOK, products)
}

func (h *ProductHandlers) GetProduct(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Product ID"})
        return
    }

    product, err := h.Service.GetProduct(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }
    c.JSON(http.StatusOK, product)
}

func (h *ProductHandlers) CreateProduct(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    if err := h.Service.CreateProduct(&product); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create product"})
        return
    }
    c.JSON(http.StatusCreated, product)
}

func (h *ProductHandlers) DeleteProduct(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Product ID"})
        return
    }

    if err := h.Service.DeleteProduct(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete product"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}