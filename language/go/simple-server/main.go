package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// CRUD API
	app.Get("/users", func(c *fiber.Ctx) error {
		return c.SendString("유저 목록 조회")
	})

	app.Get("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("유저 조회: " + id)
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		return c.SendString("유저 생성")
	})

	app.Put("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("유저 수정: " + id)
	})

	app.Delete("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("유저 삭제: " + id)
	})

	app.Listen(":3000")
}
