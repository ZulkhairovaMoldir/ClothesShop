package services

import (
	"ClothesShop/internal/models"
	"ClothesShop/internal/repository"
)

type OrderService struct {
	Repo *repository.OrderRepository
}

func (s *OrderService) CreateOrder(order *models.Order) error {
	return s.Repo.CreateOrder(order)
}

func (s *OrderService) GetOrders() ([]models.Order, error) {
	return s.Repo.GetAllOrders()
}

func (s *OrderService) GetOrder(id uint) (*models.Order, error) {
	return s.Repo.GetOrderByID(id)
}

func (s *OrderService) DeleteOrder(id uint) error {
	return s.Repo.DeleteOrder(id)
}
