package main

import (
	"gin-example/internal/handlers"
	"gin-example/internal/initializers"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
)

func init() {
	_, currentFilePathname, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(currentFilePathname)
	envPath := filepath.Join(basepath, "../../configs/.env")
	initializers.LoadEnvVariables(envPath)
	initializers.ConnectToDatabase()
}

func main() {
	router := gin.Default()

	userHandler := handlers.NewUserHandler()

	router.POST("/signup", userHandler.SignUp)

	router.Run()
}
