package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/piru72/winners-crud/server/routes"
)

func SetupApp() *fiber.App {

	app := fiber.New()

	app.Get("/api/v1/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Welcome to the Winners API!"})
	})

	routes.RegisterRoutes(app)

	return app
}
