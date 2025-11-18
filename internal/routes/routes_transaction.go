package routes

import (
	"database/sql"

	"github.com/Ntanzi07/gofinance/internal/config"
	"github.com/Ntanzi07/gofinance/internal/handlers"
	"github.com/Ntanzi07/gofinance/internal/repository"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func setupTransactionRoutes(app *fiber.App, db *sql.DB) {

	repo := repository.NewTransactionRepository(db)
	handler := handlers.NewTransactionHandler(repo)

	// Create a protected group for transactions. Only requests with a valid JWT
	// will be allowed to hit these handlers.
	protected := app.Group("/transactions")

	protected.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: config.LoadJwt()},
	}))

	// Admin endpoints for transactions
	protected.Get("/", handler.GetAllTransactionsHandler)
	protected.Get("/:id", handler.GetTransactionByIdHandler)

	// Example: route to get transactions by user (disabled/commented)
	// app.Get("/:name/transactions", handler.GetTransactionsByUserHandler)
}
