package database

import (
	"fmt"
	"log"
	"os"

	"github.com/nyshnt/go-fiber-postgres/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB represents a Database instance
var DB *gorm.DB

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

func InitDB() error {
	c := &Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMode"),
		DBName:   os.Getenv("DB_NAME"),
	}

	dns := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
		os.Exit(2)
	}

	log.Println("Database connected")

	// Register model types
	models := []interface{}{
		&models.Books{},
		&models.Users{},
	}

	log.Print("Running the migrations...")
	err = db.AutoMigrate(models...)
	if err != nil {
		log.Fatal("Error migrating models")
	}

	DB = db
	return nil
}
