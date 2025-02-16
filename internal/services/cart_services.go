package services

import (
    "ClothesShop/internal/models"
    "ClothesShop/internal/repository"
)

type CartService struct {
    Repo *repository.CartRepository
}

// Add item to cart (Works for both users and guests)
func (s *CartService) AddToCart(item *models.Cart) error {
    return s.Repo.AddItemToCart(item)
}

// Retrieve cart items for logged-in users
func (s *CartService) GetCart(customerID uint) ([]map[string]interface{}, error) {
    return s.Repo.GetCartByCustomerID(customerID)
}

// Retrieve cart items for guests
func (s *CartService) GetGuestCart(sessionID string) ([]map[string]interface{}, error) {
    return s.Repo.GetCartBySessionID(sessionID)
}

// Remove item from user cart
func (s *CartService) RemoveFromUserCart(cartItemID uint, customerID uint) error {
    return s.Repo.RemoveItemFromUserCart(cartItemID, customerID)
}