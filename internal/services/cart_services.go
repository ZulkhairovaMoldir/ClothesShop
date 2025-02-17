package services

import (
    "ClothesShop/internal/models"
    "ClothesShop/internal/repository"
)

type CartService struct {
    Repo *repository.CartRepository
}

func (s *CartService) AddToCart(cart *models.Cart) error {
    var existingCart models.Cart
    query := s.Repo.DB.Where("product_id = ?", cart.ProductID)

    if cart.CustomerID != nil {
        query = query.Where("customer_id = ?", *cart.CustomerID)
    } else if cart.SessionID != nil {
        query = query.Where("session_id = ?", *cart.SessionID)
    }

    err := query.First(&existingCart).Error
    if err == nil {
        existingCart.Quantity += cart.Quantity
        return s.Repo.DB.Save(&existingCart).Error
    }

    return s.Repo.DB.Create(cart).Error
}

func (s *CartService) UpdateCartQuantity(productID uint, newQuantity int) error {
    var cartItem models.Cart
    err := s.Repo.DB.Where("product_id = ?", productID).First(&cartItem).Error
    if err != nil {
        return err
    }

    if newQuantity <= 0 {
        return s.Repo.DB.Delete(&cartItem).Error
    }

    cartItem.Quantity = newQuantity
    return s.Repo.DB.Save(&cartItem).Error
}

func (s *CartService) RemoveFromUserCart(productID uint, customerID uint) error {
    return s.Repo.RemoveItemFromUserCart(productID, customerID)
}

func (s *CartService) RemoveFromGuestCart(productID uint, sessionID string) error {
    return s.Repo.RemoveItemFromGuestCart(productID, sessionID)
}

func (s *CartService) GetCart(customerID uint) ([]map[string]interface{}, error) {
    return s.Repo.GetCartByCustomerID(customerID)
}

func (s *CartService) GetGuestCart(sessionID string) ([]map[string]interface{}, error) {
    return s.Repo.GetCartBySessionID(sessionID)
}

func (s *CartService) MergeGuestCartToUser(sessionID string, userID uint) error {
    return s.Repo.MergeGuestCartToUser(sessionID, userID)
}