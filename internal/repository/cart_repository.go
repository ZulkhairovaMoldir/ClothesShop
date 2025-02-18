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

func (r *CartRepository) AddItemToCart(cart *models.Cart) error {
    var existingCart models.Cart
    tx := r.DB.Begin()
    defer tx.Commit()

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
        existingCart.Quantity += cart.Quantity
        return tx.Save(&existingCart).Error
    } else if errors.Is(err, gorm.ErrRecordNotFound) {
        return tx.Create(cart).Error
    } else {
        tx.Rollback()
        return err
    }
}

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
    return cartItems, err
}

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
    return cartItems, err
}

func (r *CartRepository) RemoveItemFromUserCart(cartItemID uint, customerID uint) error {
    return r.DB.Where("id = ? AND customer_id = ?", cartItemID, customerID).Delete(&models.Cart{}).Error
}

func (r *CartRepository) RemoveItemFromGuestCart(cartItemID uint, sessionID string) error {
    return r.DB.Where("id = ? AND session_id = ?", cartItemID, sessionID).Delete(&models.Cart{}).Error
}

func (r *CartRepository) MergeGuestCartToUser(sessionID string, userID uint) error {
    tx := r.DB.Begin()

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
            existingItem.Quantity += item.Quantity
            if err := tx.Save(&existingItem).Error; err != nil {
                tx.Rollback()
                return err
            }
        } else if errors.Is(err, gorm.ErrRecordNotFound) {
            item.CustomerID = &userID
            item.SessionID = nil 
            if err := tx.Save(&item).Error; err != nil {
                tx.Rollback()
                return err
            }
        } else {
            tx.Rollback()
            return err
        }
    }

    if err := tx.Where("session_id = ?", sessionID).Delete(&models.Cart{}).Error; err != nil {
        tx.Rollback()
        return err
    }

    tx.Commit()
    return nil
}