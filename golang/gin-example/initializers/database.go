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
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("database connection was not posible")
	}

	DB = db

	// process to create the custom user_type ENUM for postgres
	db.Exec(`
        DO $$ BEGIN
            CREATE TYPE user_role AS ENUM ('admin', 'customer', 'vendor');
        EXCEPTION
            WHEN duplicate_object THEN null;
        END $$;
    `)

	db.AutoMigrate(&models.User{})
}
