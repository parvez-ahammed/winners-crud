package routes

import (
	"github.com/gofiber/fiber/v2"
	v1 "github.com/piru72/winners-crud/server/routes/v1"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/api/v1/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Welcome to the Winners API!"})
	})
	v1.RegisterAllRoutes(app.Group("/api"))
}
