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

	protected := app.Group("/transactions")

	protected.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: config.LoadJwt()},
	}))

	protected.Get("/", handler.GetAllTransactionsHandler)
	protected.Get("/:id", handler.GetTransactionByIdHandler)

	//app.Get("/:name/transactions", handler.GetTransactionsByUserHandler)
}
