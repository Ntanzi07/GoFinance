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

	protected := app.Group("/user")

	protected.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: config.LoadJwt()},
	}))

	protected.Get("/:name", handler.GetUserByNameHandler)

}
