package controllers

import (
	"errors"
	"gin-example/models"
	"gin-example/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var userService = services.CreateUserService()

type UserController struct{}

type UserActions interface {
	SignUp(context *gin.Context)
}

func CreateUserController() *UserController {
	return &UserController{}
}

func (userController *UserController) SignUp(context *gin.Context) {

	var body struct {
		Name  string
		Email string
	}
	context.Bind(&body)

	user := models.User{
		ID:        uuid.NewString(),
		Name:      body.Name,
		Email:     body.Email,
		Role:      models.RoleCustomer,
		CreatedAt: time.Now(),
	}

	result := userService.CreateUser(user)
	if result.Error != nil {
		errorMessage := "There was an error, please try again later"

		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			errorMessage = "That email is already registered, please try a different one"
		}

		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": errorMessage,
		})

		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"userId": user.ID,
	})
}
