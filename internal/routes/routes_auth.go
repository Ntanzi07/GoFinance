package routes

import (
	"database/sql"

	"github.com/Ntanzi07/gofinance/internal/handlers"
	"github.com/Ntanzi07/gofinance/internal/repository"
	"github.com/gofiber/fiber/v2"
)

func setupRoutesAuth(app *fiber.App, db *sql.DB) {

	repo := repository.NewUsersRepository(db)
	handler := handlers.NewAuthHandler(repo)

	app.Post("/login", handler.LoginUserHandler)
	app.Post("/singup", handler.SingupUserHandler)

}
