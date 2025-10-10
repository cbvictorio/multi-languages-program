package controllers

import "github.com/gin-gonic/gin"

type UserController struct{}

type UserActions interface {
	SignUp(context *gin.Context)
}

func CreateUserController() *UserController {
	return &UserController{}
}

func (userController *UserController) SignUp(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "this is the user's signup endpoint",
	})
}
