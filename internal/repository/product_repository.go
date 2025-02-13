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
    result := r.DB.Table("public.products").Find(&products)
    return products, result.Error
}

func (r *ProductRepository) GetProductByID(id uint) (*models.Product, error) {
    var product models.Product
    result := r.DB.Table("public.products").First(&product, id)
    return &product, result.Error
}

func (r *ProductRepository) CreateProduct(product *models.Product) error {
    return r.DB.Table("public.products").Create(product).Error
}

func (r *ProductRepository) DeleteProduct(id uint) error {
    return r.DB.Table("public.products").Delete(&models.Product{}, id).Error
}