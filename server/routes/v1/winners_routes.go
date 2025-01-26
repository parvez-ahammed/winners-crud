package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/piru72/winners-crud/server/handlers/v1"
)

func RegisterWinnersRoutes(router fiber.Router) {
	winners := router.Group("/winners")
	winners.Get("/", handlers.GetWinners)               // GET /api/v1/winners
	winners.Get("/:id", handlers.GetWinnerByID)         // GET /api/v1/winners/:id
	winners.Post("/", handlers.CreateWinner)            // POST /api/v1/winners
	winners.Put("/:id", handlers.UpdateWinner)          // PUT /api/v1/winners/:id
	winners.Patch("/:id", handlers.PartialUpdateWinner) // PATCH /api/v1/winners/:id
	winners.Delete("/:id", handlers.DeleteWinner)       // DELETE /api/v1/winners/:id
}
