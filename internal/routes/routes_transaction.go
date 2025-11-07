package routes

import (
	"database/sql"

	"github.com/Ntanzi07/gofinance/internal/handlers"
	"github.com/Ntanzi07/gofinance/internal/repository"
	"github.com/gofiber/fiber/v2"
)

func setupTransactionRoutes(app *fiber.App, db *sql.DB) {

	repo := repository.NewTransactionRepository(db)
	handler := handlers.NewTransactionHandler(repo)

	app.Get("/transactions", handler.GetAllTransactionsHandler)
	app.Get("/transactions/:id", handler.GetTransactionByIdHandler)
	app.Post("/transactions", handler.CreateTransactionHandler)
	app.Delete("/transactions/:id", handler.DeleteTransacionHandler)

	// TODO: app.Get("/:name/transactions")
}
