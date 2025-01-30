// winner_service.go
package services

import (
	"github.com/google/uuid"
	"github.com/piru72/winners-crud/server/database/models"
	"github.com/piru72/winners-crud/server/database/repositories"
)

var winners = []models.Winner{
	{ID: "1", Season: "2024", Game: "Chess", Position: "Champion", TeamMember1: "Alice", TeamMember2: "Bob"},
	{ID: "2", Season: "2024", Game: "UNO", Position: "1st Runners Up", TeamMember1: "Bob", TeamMember2: "Charlie"},
	{ID: "3", Season: "2025", Game: "FoosBall", Position: "2nd Runners Up", TeamMember1: "Charlie", TeamMember2: "Dave"},
}

func GetWinners(season, game, position, teamMember string) ([]models.Winner, error) {
	// Get the winners from the repository
	winners, err := repositories.GetWinners(season, game, position, teamMember)
	if err != nil {
		return nil, err
	}
	return winners, nil
}

func GetWinnerByID(id string) (*models.Winner, error) {
	return repositories.GetWinnerByID(id)
}

func CreateWinner(newWinner *models.Winner) models.Winner {
	newWinner.ID = uuid.New().String()
	winners = append(winners, *newWinner)
	return *newWinner
}

func UpdateWinner(id string, updatedWinner *models.Winner) *models.Winner {
	for i, w := range winners {
		if w.ID == id {
			updatedWinner.ID = w.ID
			winners[i] = *updatedWinner
			return updatedWinner
		}
	}
	return nil
}

func PartialUpdateWinner(id string, partialWinner map[string]interface{}) *models.Winner {
	for i, w := range winners {
		if w.ID == id {
			if val, ok := partialWinner["season"]; ok {
				w.Season = val.(string)
			}
			if val, ok := partialWinner["game"]; ok {
				w.Game = val.(string)
			}
			if val, ok := partialWinner["position"]; ok {
				w.Position = val.(string)
			}
			if val, ok := partialWinner["teamMember1"]; ok {
				w.TeamMember1 = val.(string)
			}
			if val, ok := partialWinner["teamMember2"]; ok {
				w.TeamMember2 = val.(string)
			}
			winners[i] = w
			return &w
		}
	}
	return nil
}

func DeleteWinner(id string) bool {
	for i, w := range winners {
		if w.ID == id {
			winners = append(winners[:i], winners[i+1:]...)
			return true
		}
	}
	return false
}
