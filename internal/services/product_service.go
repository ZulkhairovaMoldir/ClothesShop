package services

import (
	"ClothesShop/internal/models"
	"ClothesShop/internal/repository"
)

type ProductService struct {
	Repo *repository.ProductRepository
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	return s.Repo.CreateProduct(product)
}

func (s *ProductService) GetProducts() ([]models.Product, error) {
	return s.Repo.GetAllProducts()
}

func (s *ProductService) GetProduct(id uint) (*models.Product, error) {
	return s.Repo.GetProductByID(id)
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.Repo.DeleteProduct(id)
}
