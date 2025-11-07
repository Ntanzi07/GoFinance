package routes

import (
	"github.com/Ntanzi07/gofinance/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupTransactionRoutes(app *fiber.App) {
	app.Get("/transactions", handlers.GetAllTransactionsHandler)
}
