package routes

import (
	"database/sql"

	"github.com/Ntanzi07/gofinance/internal/handlers"
	"github.com/Ntanzi07/gofinance/internal/repository"
	"github.com/gofiber/fiber/v2"
)

func setupRoutesUser(app *fiber.App, db *sql.DB) {

	repo := repository.NewUsersRepository(db)
	handler := handlers.NewUsersHandler(repo)

	app.Post("/login", handler.LoginUserHandler)

	/*
		app.Get("/users", handler.GetAllUserHandler)
		app.Get("/users/:id", handler.GetUserByIdHandler)
		app.Get("/:name", handler.GetUserByNameHandler)
		app.Post("/users", handler.CreateUserHandler)
		app.Delete("/users/:id", handler.DeleteUserHandler)
	*/
}
