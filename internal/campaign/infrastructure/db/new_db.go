package db

import (
	"fmt"
	campaign "marketplace/internal/campaign/domain"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env")
	}

	dsn := os.Getenv("DSN")
	database,err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	database.AutoMigrate(&campaign.Campaign{}, &campaign.Contact{})

	return database
}