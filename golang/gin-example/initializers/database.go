package initializers

import (
	"gin-example/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})

	if err != nil {
		log.Fatal("database connection was not posible")
	}

	DB = db

	db.Migrator().DropTable(&models.User{})
	db.AutoMigrate(&models.User{})
}
