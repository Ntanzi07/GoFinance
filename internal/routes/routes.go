package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	setupRoutesUser(app)
	setupTransactionRoutes(app)
}
