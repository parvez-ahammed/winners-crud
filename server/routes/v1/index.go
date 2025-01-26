package v1

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterAllRoutes(app fiber.Router) {
	RegisterWinnersRoutes(app.Group("/v1"))
}
