package main

import (
	"gin-example/internal/initializers"
	"gin-example/internal/models"
	"path/filepath"
	"runtime"

	log "github.com/sirupsen/logrus"
)

func init() {
	_, currentFilePathname, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(currentFilePathname)
	envPath := filepath.Join(basepath, "../configs/.env")

	initializers.LoadEnvVariables(envPath)
	initializers.ConnectToDatabase()

	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
}

func main() {
	initializers.DB.Migrator().DropTable(&models.User{})
	initializers.DB.AutoMigrate(&models.User{})
	log.Info("User Table deleted and synced successfully")
}
