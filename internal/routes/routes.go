package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *sql.DB) {

	setupRoutesAuth(app, db)
	setupTransactionRoutes(app, db)
	setupRoutesUser(app, db)
}
