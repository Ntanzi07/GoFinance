package routes

import (
	"database/sql"

	_ "github.com/Ntanzi07/gofinance/docs"
	fiberSwagger "github.com/swaggo/fiber-swagger"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *sql.DB) {

	app.Get("/docs/*", fiberSwagger.WrapHandler)

	setupRoutesAuth(app, db)
	setupTransactionRoutes(app, db)
	setupRoutesUser(app, db)
}
