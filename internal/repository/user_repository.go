package repositories

import (
	"ClothesShop/config"
	"ClothesShop/internal/models"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := config.DB.Find(&users)
	return users, result.Error
}

func GetUserByID(id uint) (models.User, error) {
	var user models.User
	result := config.DB.First(&user, id)
	return user, result.Error
}

func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

func DeleteUser(id uint) error {
	return config.DB.Delete(&models.User{}, id).Error
}
