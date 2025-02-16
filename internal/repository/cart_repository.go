package repository

import (
    "ClothesShop/internal/models"
    "gorm.io/gorm"
)

type CartRepository struct {
    DB *gorm.DB
}

// Add an item to the cart
func (r *CartRepository) AddItemToCart(cart *models.Cart) error {
    var existingCart models.Cart

    if cart.CustomerID != nil {
        // Check if product already exists in cart
        err := r.DB.Where("customer_id = ? AND product_id = ?", *cart.CustomerID, cart.ProductID).First(&existingCart).Error
        if err == nil {
            // If product exists, update quantity instead of inserting a new row
            existingCart.Quantity += cart.Quantity
            return r.DB.Save(&existingCart).Error
        } else if err != gorm.ErrRecordNotFound {
            return err // Return DB error if it's not "record not found"
        }
    } else if cart.SessionID != nil {
        // Check if product already exists in guest's cart
        err := r.DB.Where("session_id = ? AND product_id = ?", *cart.SessionID, cart.ProductID).First(&existingCart).Error
        if err == nil {
            // If product exists, update quantity instead of inserting a new row
            existingCart.Quantity += cart.Quantity
            return r.DB.Save(&existingCart).Error
        } else if err != gorm.ErrRecordNotFound {
            return err // Return DB error if it's not "record not found"
        }
    }

    // If product is not in the cart, insert a new row
    return r.DB.Create(cart).Error
}

// Retrieve cart items by Customer ID (for logged-in users)
func (r *CartRepository) GetCartByCustomerID(customerID uint) ([]map[string]interface{}, error) {
    var cartItems []map[string]interface{}

    query := `
        SELECT c.id, c.product_id, p.name, p.description, p.price, c.quantity
        FROM public.carts c
        JOIN public.products p ON c.product_id = p.id
        WHERE c.customer_id = ?
    `

    err := r.DB.Raw(query, customerID).Scan(&cartItems).Error
    return cartItems, err
}

// Retrieve cart items by Session ID (for guest users)
func (r *CartRepository) GetCartBySessionID(sessionID string) ([]map[string]interface{}, error) {
    var cartItems []map[string]interface{}

    query := `
        SELECT c.id, c.product_id, p.name, p.description, p.price, c.quantity
        FROM public.carts c
        JOIN public.products p ON c.product_id = p.id
        WHERE c.session_id = ?
    `

    err := r.DB.Raw(query, sessionID).Scan(&cartItems).Error
    return cartItems, err
}

// Remove item from user cart
func (r *CartRepository) RemoveItemFromUserCart(cartItemID uint, customerID uint) error {
    return r.DB.Where("id = ? AND customer_id = ?", cartItemID, customerID).Delete(&models.Cart{}).Error
}