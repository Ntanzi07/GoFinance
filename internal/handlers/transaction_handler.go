package handlers

import "github.com/gofiber/fiber/v2"

func GetAllTransactionsHandler(c *fiber.Ctx) error {
	return c.SendString("Get all transactions")
}
