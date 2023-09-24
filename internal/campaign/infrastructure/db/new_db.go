package db

import (
	campaign "marketplace/internal/campaign/domain"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := os.Getenv("DSN")
	database,err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	database.AutoMigrate(&campaign.Campaign{}, &campaign.Contact{})

	return database
}