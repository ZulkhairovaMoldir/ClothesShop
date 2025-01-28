package services

import (
	"ClothesShop/internal/models"
	"ClothesShop/internal/repository"
)

type CartService struct {
	Repo *repository.CartRepository
}

func (s *CartService) AddToCart(item *models.Cart) error {
	return s.Repo.AddItemToCart(item)
}

func (s *CartService) GetCart(customerID uint) ([]models.Cart, error) {
	return s.Repo.GetCartByCustomerID(customerID)
}

func (s *CartService) RemoveFromCart(cartID uint) error {
	return s.Repo.RemoveItemFromCart(cartID)
}