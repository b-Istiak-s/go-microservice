package db

import (
	"fmt"
	"log"
	"myapp/internal/model"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dotEnvError := godotenv.Load()
	if dotEnvError != nil {
		log.Fatal("Error loading .env file")
	}
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_SSL_MODE"),
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
