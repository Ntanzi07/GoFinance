package routes

import (
	"database/sql"

	"github.com/Ntanzi07/gofinance/internal/handlers"
	"github.com/Ntanzi07/gofinance/internal/repository"
	"github.com/gofiber/fiber/v2"
)

func setupRoutesAuth(app *fiber.App, db *sql.DB) {
	// Create repository and handler instances and wire them to endpoints.
	// The repository receives the shared *sql.DB; the handler receives the
	// repository (dependency injection). These endpoints are intentionally
	// public so clients can obtain tokens and register new users.
	repo := repository.NewUsersRepository(db)
	handler := handlers.NewAuthHandler(repo)

	app.Post("/login", handler.LoginUserHandler)
	app.Post("/singup", handler.SingupUserHandler)

}
