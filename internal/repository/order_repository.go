package repository

import (
	"ClothesShop/internal/models"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func (r *OrderRepository) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	result := r.DB.Find(&orders)
	return orders, result.Error
}

func (r *OrderRepository) GetOrderByID(id uint) (*models.Order, error) {
	var order models.Order
	result := r.DB.First(&order, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

func (r *OrderRepository) CreateOrder(order *models.Order) error {
	return r.DB.Create(order).Error
}

func (r *OrderRepository) DeleteOrder(id uint) error {
	return r.DB.Delete(&models.Order{}, id).Error
}
