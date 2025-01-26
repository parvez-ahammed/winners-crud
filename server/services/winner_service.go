// winner_service.go
package services

import (
	"github.com/google/uuid"
	"github.com/piru72/winners-crud/server/database/entities"
)

var winners = []entities.Winner{
	{ID: "1", Season: "2024", Game: "Chess", Position: "Champion", TeamMember1: "Alice", TeamMember2: "Bob"},
	{ID: "2", Season: "2024", Game: "UNO", Position: "1st Runners Up", TeamMember1: "Bob", TeamMember2: "Charlie"},
	{ID: "3", Season: "2025", Game: "FoosBall", Position: "2nd Runners Up", TeamMember1: "Charlie", TeamMember2: "Dave"},
}

func GetWinnersService(season, game, position, teamMember string) []entities.Winner {
	var filtered []entities.Winner
	for _, w := range winners {
		if (season == "" || w.Season == season) &&
			(game == "" || w.Game == game) &&
			(position == "" || w.Position == position) &&
			(teamMember == "" || w.TeamMember1 == teamMember || w.TeamMember2 == teamMember) {
			filtered = append(filtered, w)
		}
	}
	return filtered
}

func GetWinnerByIDService(id string) *entities.Winner {
	for _, w := range winners {
		if w.ID == id {
			return &w
		}
	}
	return nil
}

func CreateWinnerService(newWinner *entities.Winner) entities.Winner {
	newWinner.ID = uuid.New().String()
	winners = append(winners, *newWinner)
	return *newWinner
}

func UpdateWinnerService(id string, updatedWinner *entities.Winner) *entities.Winner {
	for i, w := range winners {
		if w.ID == id {
			updatedWinner.ID = w.ID
			winners[i] = *updatedWinner
			return updatedWinner
		}
	}
	return nil
}

func PartialUpdateWinnerService(id string, partialWinner map[string]interface{}) *entities.Winner {
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

func DeleteWinnerService(id string) bool {
	for i, w := range winners {
		if w.ID == id {
			winners = append(winners[:i], winners[i+1:]...)
			return true
		}
	}
	return false
}
