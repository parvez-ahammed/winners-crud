package database

import (
	"log"

	"github.com/piru72/winners-crud/server/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	dsn := "host=localhost user=user1 password=password1 dbname=winner_crud port=5432 sslmode=disable TimeZone=Asia/Shanghai"

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
