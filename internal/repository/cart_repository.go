package repository

import (
    "ClothesShop/internal/models"
    "errors"
    "gorm.io/gorm"
    "log"
)

type CartRepository struct {
    DB *gorm.DB
}

// Add an item to the cart
func (r *CartRepository) AddItemToCart(cart *models.Cart) error {
    var existingCart models.Cart
    tx := r.DB.Begin()
    defer tx.Commit()

    // Ensure query includes correct session or customer ID
    query := tx.Where("product_id = ?", cart.ProductID)
    if cart.CustomerID != nil {
        query = query.Where("customer_id = ?", *cart.CustomerID)
        log.Printf("Adding to cart for customer_id: %d, product_id: %d", *cart.CustomerID, cart.ProductID)
    } else if cart.SessionID != nil {
        query = query.Where("session_id = ?", *cart.SessionID)
        log.Printf("Adding to cart for session_id: %s, product_id: %d", *cart.SessionID, cart.ProductID)
    } else {
        tx.Rollback()
        return errors.New("either CustomerID or SessionID must be provided")
    }

    err := query.First(&existingCart).Error
    if err == nil {
        // Product exists, update quantity properly
        existingCart.Quantity += cart.Quantity
        return tx.Save(&existingCart).Error
    } else if errors.Is(err, gorm.ErrRecordNotFound) {
        // Product does not exist, insert as a new row
        return tx.Create(cart).Error
    } else {
        tx.Rollback()
        return err
    }
}

// Retrieve cart items by Customer ID (for logged-in users)
func (r *CartRepository) GetCartByCustomerID(customerID uint) ([]map[string]interface{}, error) {
    var cartItems []map[string]interface{}

    query := `
        SELECT c.product_id, p.name, p.description, p.price, SUM(c.quantity) AS quantity
        FROM public.carts c
        JOIN public.products p ON c.product_id = p.id
        WHERE c.customer_id = ?
        GROUP BY c.product_id, p.name, p.description, p.price
    `

    err := r.DB.Raw(query, customerID).Scan(&cartItems).Error
    if cartItems == nil {
        cartItems = []map[string]interface{}{}
    }
    return cartItems, err
}

// Retrieve cart items by Session ID (for guest users)
func (r *CartRepository) GetCartBySessionID(sessionID string) ([]map[string]interface{}, error) {
    var cartItems []map[string]interface{}

    query := `
        SELECT c.product_id, p.name, p.description, p.price, SUM(c.quantity) AS quantity
        FROM public.carts c
        JOIN public.products p ON c.product_id = p.id
        WHERE c.session_id = ?
        GROUP BY c.product_id, p.name, p.description, p.price
    `

    err := r.DB.Raw(query, sessionID).Scan(&cartItems).Error
    if cartItems == nil {
        cartItems = []map[string]interface{}{}
    }
    return cartItems, err
}

// Remove item from user cart
func (r *CartRepository) RemoveItemFromUserCart(cartItemID uint, customerID uint) error {
    return r.DB.Where("id = ? AND customer_id = ?", cartItemID, customerID).Delete(&models.Cart{}).Error
}

// Remove item from guest cart
func (r *CartRepository) RemoveItemFromGuestCart(cartItemID uint, sessionID string) error {
    return r.DB.Where("id = ? AND session_id = ?", cartItemID, sessionID).Delete(&models.Cart{}).Error
}

// Transfer guest cart to user cart
func (r *CartRepository) MergeGuestCartToUser(sessionID string, userID uint) error {
    tx := r.DB.Begin()

    // Step 1: Check if guest cart has items
    var guestCartItems []models.Cart
    err := tx.Where("session_id = ?", sessionID).Find(&guestCartItems).Error
    if err != nil {
        tx.Rollback()
        return err
    }

    for _, item := range guestCartItems {
        var existingItem models.Cart
        query := tx.Where("product_id = ? AND customer_id = ?", item.ProductID, userID)

        err := query.First(&existingItem).Error
        if err == nil {
            // If item exists in user's cart, update quantity
            existingItem.Quantity += item.Quantity
            if err := tx.Save(&existingItem).Error; err != nil {
                tx.Rollback()
                return err
            }
        } else if errors.Is(err, gorm.ErrRecordNotFound) {
            // If item does not exist, assign it to the user
            item.CustomerID = &userID
            item.SessionID = nil // Remove session ID
            if err := tx.Save(&item).Error; err != nil {
                tx.Rollback()
                return err
            }
        } else {
            tx.Rollback()
            return err
        }
    }

    // Step 3: Delete old session cart records
    if err := tx.Where("session_id = ?", sessionID).Delete(&models.Cart{}).Error; err != nil {
        tx.Rollback()
        return err
    }

    tx.Commit()
    return nil
}