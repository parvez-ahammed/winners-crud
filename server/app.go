package main

import (
	"github.com/gofiber/fiber/v2"
	database "github.com/piru72/winners-crud/server/config"
	"github.com/piru72/winners-crud/server/routes"
)

func SetupApp() *fiber.App {

	// Initialize the database
	database.Connect()

	// Run Migrations
	//database.RunMigrations(db)

	// Seed Data
	// winner.SeedWinner(db)
	app := fiber.New()

	app.Get("/api/v1/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Welcome to the Winners API!"})
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Welcome to the Winners API ROOT! "})
	})

	routes.RegisterRoutes(app)

	return app
}
