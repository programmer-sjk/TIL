package handlers

import "github.com/gofiber/fiber/v2"

func GetUsers(c *fiber.Ctx) error {
	return c.SendString("유저 목록 조회")
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("유저 조회: " + id)
}

func CreateUser(c *fiber.Ctx) error {
	return c.SendString("유저 생성")
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("유저 수정: " + id)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("유저 삭제: " + id)
}
