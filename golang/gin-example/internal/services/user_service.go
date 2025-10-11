package services

import (
	"errors"
	"fmt"
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
		return nil, errors.New("there was an error retrieving all users")
	}

	return users, nil
}

func (userService *UserService) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := initializers.DB.Where(&models.User{Email: email}).Find(&user)
	userNotFoundError := fmt.Sprintf("there was an error retrieving the user with email: %s", email)

	if result.Error != nil {
		return nil, errors.New(userNotFoundError)
	}

	return &user, nil
}
