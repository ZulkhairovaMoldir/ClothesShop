package services

import (
	"ClothesShop/internal/models"
	"ClothesShop/internal/repository"
	"errors"
)

type UserService struct {
	Repo *repository.UserRepository
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.Repo.Save(user)
}

func (s *UserService) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := s.Repo.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.Repo.GetAllUsers()
}

func (s *UserService) GetUserByID(id string) (*models.User, error) {
	return s.Repo.GetUserByID(id)
}

func (s *UserService) DeleteUser(id string) error {
	return s.Repo.DeleteUser(id)
}
