package db

import (
	"fmt"
	"log"
	"myapp/internal/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)
	// Connect to the database
	// Use the DSN to connect to the PostgreSQL database
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Auto migrate models
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("Error performing database migrations: %v", err)
	}

	return err
}

func CloseDB() {
	sqlDB, err := DB.DB() // Get the underlying sql.DB instance
	if err != nil {
		log.Printf("Error retrieving database instance: %v", err)
		return
	}

	if err := sqlDB.Close(); err != nil {
		log.Printf("Error closing the database: %v", err)
	}
}
