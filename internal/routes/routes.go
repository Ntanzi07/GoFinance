package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *sql.DB) {

	setupRoutesUser(app, db)
	setupTransactionRoutes(app, db)
}
