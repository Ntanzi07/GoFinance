package routes

import (
	"database/sql"

	"github.com/Ntanzi07/gofinance/internal/config"
	"github.com/Ntanzi07/gofinance/internal/handlers"
	"github.com/Ntanzi07/gofinance/internal/repository"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func setupRoutesUser(app *fiber.App, db *sql.DB) {

	repo := repository.NewUsersRepository(db)
	handler := handlers.NewUsersHandler(repo)

	// Group routes by username. All routes inside this group require a valid JWT.
	protected := app.Group("/:name")

	// Apply JWT middleware to the group using the secret from config
	protected.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: config.LoadJwt()},
	}))

	// User info and transaction operations scoped to the route parameter `name`.
	protected.Get("/infos", handler.GetUserByNameHandler)
	protected.Get("/", handler.GetUserTransactions)
	protected.Post("/", handler.CreateUserTransaction)
	protected.Put("/transactions/:id", handler.UpdateUserTransaction)
	protected.Delete("/transactions/:id", handler.DeleteUserTransaction)
}
