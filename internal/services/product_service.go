package services

import (
	"ClothesShop/internal/models"
	"ClothesShop/internal/repository"
	"log"
)

type ProductService struct {
	Repo *repository.ProductRepository
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	return s.Repo.CreateProduct(product)
}

func (s *ProductService) GetProducts() ([]models.Product, error) {
	var products []models.Product
	log.Println("Fetching products from database...") // Debug log
	result := s.Repo.DB.Find(&products)
	if result.Error != nil {
		log.Println("Database query error:", result.Error)
		return nil, result.Error
	}
	if len(products) == 0 {
		log.Println("Warning: No products found in the database!")
	} else {
		log.Println("Products from DB:", products) // Debugging log
	}
	return products, nil
}

func (s *ProductService) GetProduct(id uint) (*models.Product, error) {
	return s.Repo.GetProductByID(id)
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.Repo.DeleteProduct(id)
}
