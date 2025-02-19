package handlers

import (
    "ClothesShop/internal/models"
    "ClothesShop/internal/services"
    "ClothesShop/internal/utils"
    "ClothesShop/middleware"
    "github.com/gin-contrib/sessions"
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
    "strconv"
)

type UserHandlers struct {
    Service *services.UserService
    CartService *services.CartService
}

type AuthHandler struct {
    Service *services.UserService
    CartService *services.CartService 
}

func (h *AuthHandler) Login(c *gin.Context) {
    var req struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    user, err := h.Service.FindByEmail(req.Email)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email"})
        return
    }

    if err := utils.CheckPasswordHash(req.Password, user.Password); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
        return
    }

    token, err := middleware.GenerateJWT(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    session := sessions.Default(c)
    sessionID, sessionExists := session.Get("sessionID").(string)

    if sessionExists && sessionID != "" {
        err := h.CartService.MergeGuestCartToUser(sessionID, user.ID)
        if err != nil {
            log.Printf("Failed to merge guest cart: %v", err)
        }
        session.Delete("sessionID") 
        session.Save()
    }

    session.Set("customerID", user.ID)
    session.Save()

    c.Set("customerID", user.ID)
    c.JSON(http.StatusOK, gin.H{"token": token, "message": "Login successful"})
}

func (h *UserHandlers) CreateUser(c *gin.Context) {
    var req struct {
        Name     string `json:"name" binding:"required"`
        Email    string `json:"email" binding:"required"`
        Password string `json:"password" binding:"required"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    hashedPassword, err := utils.HashPassword(req.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }

    user := &models.User{
        Name:     req.Name,
        Email:    req.Email,
        Password: hashedPassword,
    }
    if err := h.Service.CreateUser(user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (h *UserHandlers) GetUsers(c *gin.Context) {
    users, err := h.Service.GetAllUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
        return
    }
    c.JSON(http.StatusOK, users)
}

func (h *UserHandlers) GetUser(c *gin.Context) {
    id := c.Param("id")
    user, err := h.Service.GetUserByID(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
        return
    }
    c.JSON(http.StatusOK, user)
}

func (h *UserHandlers) GetProfile(c *gin.Context) {
    customerID, exists := c.Get("customerID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    customerIDStr := strconv.Itoa(int(customerID.(uint)))

    user, err := h.Service.GetUserByID(customerIDStr)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "name":  user.Name,
        "email": user.Email,
    })
}

func (h *UserHandlers) DeleteUser(c *gin.Context) {
    id := c.Param("id")
    err := h.Service.DeleteUser(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}