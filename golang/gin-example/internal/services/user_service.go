package services

import (
	"errors"
	"gin-example/internal/initializers"
	"gin-example/internal/models"

	"gorm.io/gorm"
)

type UserService struct{}

type UserHandler interface {
	CreateUser(user models.User) *gorm.DB
	GetAllUsers(user models.User) *gorm.DB
}

func NewUserService() *UserService {
	return &UserService{}
}

func (userService *UserService) CreateUser(user models.User) error {
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		errorMessage := "There was an error, please try again later"

		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			errorMessage = "That email is already registered, please try a different one"
		}

		return errors.New(errorMessage)
	}

	return result.Error
}

func (userService *UserService) GetAllUsers(user models.User) ([]models.User, error) {
	var users []models.User
	result := initializers.DB.Find(&user)

	if result.Error != nil {
		errorMsg := "There was an error retrieving all users"
		return nil, errors.New(errorMsg)
	}

	return users, nil
}
