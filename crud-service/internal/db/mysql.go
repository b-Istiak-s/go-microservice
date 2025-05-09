package db

import (
	"crud-service/internal/model"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dotEnvError := godotenv.Load()
	if dotEnvError != nil {
		log.Fatal("Error loading .env file")
	}
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DB"),
	)
	// Connect to the database
	// Use the DSN to connect to the PostgreSQL database
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Auto migrate models
	err = DB.AutoMigrate(&model.Note{})
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
