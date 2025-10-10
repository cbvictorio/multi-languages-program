package main

import (
	"gin-example/controllers"
	"gin-example/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	router := gin.Default()

	userController := controllers.CreateUserController()

	router.GET("/signup", userController.SignUp)

	router.Run()
}
