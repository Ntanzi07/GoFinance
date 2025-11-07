package routes

import (
	"github.com/Ntanzi07/gofinance/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutesUser(app *fiber.App) {
	app.Get("/users", handlers.GetAllUserHandler)
}
