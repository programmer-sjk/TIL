package main

import (
	"go-fiber-tutorial/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	routes.SetupUserRoutes(app)

	app.Listen(":3000")
}
