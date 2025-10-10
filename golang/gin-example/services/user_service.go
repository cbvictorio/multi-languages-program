package services

import (
	"gin-example/initializers"
	"gin-example/models"

	"gorm.io/gorm"
)

type UserService struct{}

type UserHandler interface {
	CreateUser(user models.User) *gorm.DB
	GetAllUsers(user models.User) *gorm.DB
}

func CreateUserService() *UserService {
	return &UserService{}
}

func (userService *UserService) CreateUser(user models.User) *gorm.DB {
	result := initializers.DB.Create(&user)
	return result
}

func (userService *UserService) GetAllUsers(user models.User) *gorm.DB {
	result := initializers.DB.Create(&user)
	return result
}
