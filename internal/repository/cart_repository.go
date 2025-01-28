package repository

import (
	"ClothesShop/internal/models"
	"gorm.io/gorm"
)

type CartRepository struct {
	DB *gorm.DB
}

func (r *CartRepository) GetCartByCustomerID(customerID uint) ([]models.Cart, error) {
	var cartItems []models.Cart
	result := r.DB.Where("customer_id = ?", customerID).Find(&cartItems)
	return cartItems, result.Error
}

func (r *CartRepository) AddItemToCart(cart *models.Cart) error {
	return r.DB.Create(cart).Error
}

func (r *CartRepository) UpdateCartItem(cart *models.Cart) error {
	return r.DB.Save(cart).Error
}

func (r *CartRepository) RemoveItemFromCart(cartID uint) error {
	return r.DB.Delete(&models.Cart{}, cartID).Error
}
