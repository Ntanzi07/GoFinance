package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func setupRoutesUser(app *fiber.App, db *sql.DB) {

	//repo := repository.NewUsersRepository(db)
	//handler := handlers.NewUsersHandler(repo)

	//app.Get("/:name", handler.GetUserByNameHandler)

}
