package handlers

import (
	"gin-example/internal/models"
	"gin-example/internal/services"
	"gin-example/pkg/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	INCORRECT_EMAIL_OR_PASSWORD_MSG = "The email or password are incorrect, please try again"
)

var userService = services.NewUserService()

type UserHandler struct{}
type UserActions interface {
	SignUp(context *gin.Context)
	Login(context *gin.Context)
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (userHandler *UserHandler) SignUp(context *gin.Context) {

	var body struct {
		Name     string
		Email    string
		Password string
	}
	context.Bind(&body)

	password, err := utils.GenerateHashFromPassword(body.Password)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	user := models.User{
		ID:        uuid.NewString(),
		Name:      body.Name,
		Email:     body.Email,
		Role:      models.RoleCustomer,
		CreatedAt: time.Now(),
		Password:  password,
	}

	err = userService.CreateUser(user)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "user created",
	})
}

func (userHandler *UserHandler) Login(context *gin.Context) {
	var body struct {
		Email    string
		Password string
	}
	context.Bind(&body)

	user, err := userService.GetUserByEmail(body.Email)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	if user == nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": INCORRECT_EMAIL_OR_PASSWORD_MSG,
		})

		return
	}

	isCorrectPassword := utils.CompareHashWithPassword(user.Password, body.Password)
	if !isCorrectPassword {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": INCORRECT_EMAIL_OR_PASSWORD_MSG,
		})

		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "login successful",
	})
}
