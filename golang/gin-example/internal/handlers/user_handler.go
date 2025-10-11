package handlers

import (
	"gin-example/internal/models"
	"gin-example/internal/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var userService = services.NewUserService()

type UserHandler struct{}

type UserActions interface {
	SignUp(context *gin.Context)
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (userHandler *UserHandler) SignUp(context *gin.Context) {

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

	err := userService.CreateUser(user)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	context.JSON(http.StatusCreated, user)
}
