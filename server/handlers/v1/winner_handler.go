package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/piru72/winners-crud/server/database/models"
	"github.com/piru72/winners-crud/server/services"
	"github.com/piru72/winners-crud/server/utils"
)

func GetWinners(c *fiber.Ctx) error {
	season := c.Query("season")
	game := c.Query("game")
	position := c.Query("position")
	teamMember := c.Query("teamMember")
	filtered, err := services.GetWinners(season, game, position, teamMember)
	if err != nil {
		return utils.SendApiResponse(c, false, 500, "Error fetching winners", nil, err.Error())
	}

	return utils.SendApiResponse(c, true, 200, "Winners fetched successfully", filtered, nil)
}

func GetWinnerByID(c *fiber.Ctx) error {
	id := c.Params("id")
	winner, err := services.GetWinnerByID(id)
	if err != nil {
		return utils.SendApiResponse(c, false, 404, "Winner not found", nil, err.Error())
	}
	return utils.SendApiResponse(c, true, 200, "Winner fetched successfully", winner, nil)
}

func CreateWinner(c *fiber.Ctx) error {
	var newWinner models.Winner
	if err := c.BodyParser(&newWinner); err != nil {
		return utils.SendApiResponse(c, false, 400, "Invalid input", nil, err.Error())
	}

	createdWinner := services.CreateWinner(&newWinner)
	return utils.SendApiResponse(c, true, 201, "Winner created successfully", createdWinner, nil)
}

func UpdateWinner(c *fiber.Ctx) error {
	id := c.Params("id")
	var updatedWinner models.Winner
	if err := c.BodyParser(&updatedWinner); err != nil {
		return utils.SendApiResponse(c, false, 400, "Invalid input", nil, err.Error())
	}

	updated := services.UpdateWinner(id, &updatedWinner)
	if updated == nil {
		return utils.SendApiResponse(c, false, 404, "Winner not found", nil, "Winner not found")
	}

	return utils.SendApiResponse(c, true, 200, "Winner updated successfully", updated, nil)
}

func PartialUpdateWinner(c *fiber.Ctx) error {
	id := c.Params("id")
	var partialWinner map[string]interface{}
	if err := c.BodyParser(&partialWinner); err != nil {
		return utils.SendApiResponse(c, false, 400, "Invalid input", nil, err.Error())
	}

	updated := services.PartialUpdateWinner(id, partialWinner)
	if updated == nil {
		return utils.SendApiResponse(c, false, 404, "Winner not found", nil, "Winner not found")
	}
	return utils.SendApiResponse(c, true, 200, "Winner partially updated successfully", updated, nil)
}

func DeleteWinner(c *fiber.Ctx) error {
	id := c.Params("id")
	success := services.DeleteWinner(id)
	if !success {
		return utils.SendApiResponse(c, false, 404, "Winner not found", nil, "Winner not found")
	}
	return utils.SendApiResponse(c, true, 204, "Winner deleted successfully", nil, nil)
}
