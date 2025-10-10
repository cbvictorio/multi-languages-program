package main

import (
	"gin-example/initializers"
	"gin-example/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
