package repositories

import (
	database "github.com/piru72/winners-crud/server/config"
	"github.com/piru72/winners-crud/server/database/models"
)

func GetWinners(season, game, position, teamMember string) ([]models.Winner, error) {
	var winners []models.Winner
	query := database.DB.Model(&models.Winner{})

	// Apply filters
	if season != "" {
		query = query.Where("season = ?", season)
	}
	if game != "" {
		query = query.Where("game = ?", game)
	}
	if position != "" {
		query = query.Where("position = ?", position)
	}
	if teamMember != "" {
		query = query.Where("team_member1 = ? OR team_member2 = ?", teamMember, teamMember)
	}

	// Execute the query
	if err := query.Find(&winners).Error; err != nil {
		return nil, err
	}

	return winners, nil
}
func GetWinnerByID(id string) (*models.Winner, error) {
	var winner models.Winner
	// Query the database for the winner by ID
	if err := database.DB.Where("id = ?", id).First(&winner).Error; err != nil {
		return nil, err
	}
	return &winner, nil
}
