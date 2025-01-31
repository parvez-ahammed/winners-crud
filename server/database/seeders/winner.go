package seeders

import (
	"log"

	"github.com/piru72/winners-crud/server/database/models"
	"gorm.io/gorm"
)

func SeedWinner(db *gorm.DB) {
	var count int64
	db.Model(&models.Winner{}).Count(&count)

	// Avoid duplicate seeding
	if count == 0 {
		winners := []models.Winner{
			{Season: "Spring 2024", Game: "Dota 2", Position: "1st", TeamMember1: "PlayerA", TeamMember2: "PlayerB"},
			{Season: "Summer 2024", Game: "CS:GO", Position: "2nd", TeamMember1: "PlayerC", TeamMember2: "PlayerD"},
			{Season: "Fall 2024", Game: "Valorant", Position: "3rd", TeamMember1: "PlayerE", TeamMember2: "PlayerF"},
		}

		if err := db.Create(&winners).Error; err != nil {
			log.Fatalf("Error seeding data: %v", err)
		}

		log.Println("Seed data inserted successfully")
	} else {
		log.Println("Database already seeded")
	}
}
