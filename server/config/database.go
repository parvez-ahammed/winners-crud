package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/piru72/winners-crud/server/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	// Load environment variables from .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Warning: No .env file found. Using system environment variables.")
	}

	// Read database config from environment variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslMode := os.Getenv("DB_SSLMODE")
	timeZone := os.Getenv("DB_TIMEZONE")

	// Construct DSN
	dsn := "host=" + host +
		" user=" + user +
		" password=" + password +
		" dbname=" + dbName +
		" port=" + port +
		" sslmode=" + sslMode +
		" TimeZone=" + timeZone

	// Connect to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = db
	log.Println("Database connected successfully")
}

// RunMigrations creates the necessary tables
func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Winner{}, // Add all models here
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Database migrated successfully")
}
