package handlers

import "github.com/gofiber/fiber/v2"

func GetAllUserHandler(c *fiber.Ctx) error {
	return c.SendString("Get all users")
}
