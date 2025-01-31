package services

import (
	"github.com/piru72/winners-crud/server/database/models"
	"github.com/piru72/winners-crud/server/database/repositories"
	"github.com/pkg/errors"
)

func GetWinners(season, game, position, teamMember string) ([]models.Winner, error) {
	winners, err := repositories.GetWinners(season, game, position, teamMember)
	if err != nil {
		return nil, err
	}
	return winners, nil
}

func GetWinnerByID(id string) (*models.Winner, error) {
	return repositories.GetWinnerByID(id)
}

func CreateWinner(winner *models.Winner) error {
	return repositories.CreateWinner(winner)
}

func UpdateWinner(id string, updatedWinner *models.Winner) (*models.Winner, error) {
	existingWinner, err := repositories.GetWinnerByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "winner not found")
	}

	existingWinner.Season = updatedWinner.Season
	existingWinner.Game = updatedWinner.Game
	existingWinner.Position = updatedWinner.Position
	existingWinner.TeamMember1 = updatedWinner.TeamMember1
	existingWinner.TeamMember2 = updatedWinner.TeamMember2

	updated, err := repositories.UpdateWinner(id, existingWinner)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update winner")
	}

	return updated, nil
}

func PartialUpdateWinner(id string, partialWinner map[string]interface{}) (*models.Winner, error) {
	existingWinner, err := repositories.GetWinnerByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "winner not found")
	}

	if season, ok := partialWinner["season"]; ok {
		existingWinner.Season = season.(string)
	}
	if game, ok := partialWinner["game"]; ok {
		existingWinner.Game = game.(string)
	}
	if position, ok := partialWinner["position"]; ok {
		existingWinner.Position = position.(string)
	}
	if teamMember1, ok := partialWinner["teamMember1"]; ok {
		existingWinner.TeamMember1 = teamMember1.(string)
	}
	if teamMember2, ok := partialWinner["teamMember2"]; ok {
		existingWinner.TeamMember2 = teamMember2.(string)
	}

	updated, err := repositories.UpdateWinner(id, existingWinner)
	if err != nil {
		return nil, errors.Wrap(err, "failed to partially update winner")
	}

	return updated, nil
}

func DeleteWinner(id string) (bool, error) {
	_, err := repositories.GetWinnerByID(id)
	if err != nil {
		return false, errors.Wrap(err, "winner not found")
	}

	err = repositories.DeleteWinner(id)
	if err != nil {
		return false, errors.Wrap(err, "failed to delete winner")
	}

	return true, nil
}
