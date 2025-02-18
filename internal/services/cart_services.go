package services

import (
    "ClothesShop/internal/models"
    "ClothesShop/internal/repository"
)

type CartService struct {
    Repo *repository.CartRepository
}

func NewCartService(repo *repository.CartRepository) *CartService {
    return &CartService{Repo: repo}
}

func (s *CartService) AddToCart(cart *models.Cart) error {
    return s.Repo.AddItemToCart(cart)
}

func (s *CartService) UpdateCartQuantity(productID uint, newQuantity int, sessionID *string, customerID *uint) (*models.Cart, error) {
    var cartItem models.Cart
    query := s.Repo.DB.Where("product_id = ?", productID)

    if customerID != nil {
        query = query.Where("customer_id = ?", *customerID)
    } else if sessionID != nil {
        query = query.Where("session_id = ?", *sessionID)
    }

    err := query.First(&cartItem).Error
    if err != nil {
        return nil, err
    }

    if newQuantity <= 0 {
        s.Repo.DB.Delete(&cartItem)
        return nil, nil
    }

    cartItem.Quantity = newQuantity
    err = s.Repo.DB.Save(&cartItem).Error
    if err != nil {
        return nil, err
    }

    return &cartItem, nil
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