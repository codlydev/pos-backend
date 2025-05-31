package config

import (
	"log"
	"os"
	"pos-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failedto connect to database:", err)
	}
	//  Auto-migrate tables
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Sale{})
	DB = db
}
