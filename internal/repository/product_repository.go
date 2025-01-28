package repository

import (
	"ClothesShop/internal/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func (r *ProductRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	result := r.DB.Find(&products)
	return products, result.Error
}

func (r *ProductRepository) GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	result := r.DB.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (r *ProductRepository) CreateProduct(product *models.Product) error {
	return r.DB.Create(product).Error
}

func (r *ProductRepository) DeleteProduct(id uint) error {
	return r.DB.Delete(&models.Product{}, id).Error
}

func (r *ProductRepository) UpdateProduct(product *models.Product) error {
	return r.DB.Save(product).Error
}
