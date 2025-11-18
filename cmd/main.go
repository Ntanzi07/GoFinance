package main

import (
	"github.com/Ntanzi07/gofinance/internal/database"
	"github.com/Ntanzi07/gofinance/internal/routes"
	"github.com/gofiber/fiber/v2"
)

// @title API Fiber
// @version 1.0
// @description Documentação da API GoFinance.
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// Initialize database connection. Connect() returns a *sql.DB which is
	// safe for concurrent use across handlers and repositories. We defer
	// closing it so resources are released when the process exits.
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create a Fiber application instance. Configure middleware here if
	// needed (logging, CORS, recover, etc.).
	app := fiber.New()

	// Register routes, passing the shared DB connection so repositories can
	// be constructed with dependency injection.
	routes.SetupRoutes(app, db)

	// Start the HTTP server. Use APP_PORT from env in the future instead of
	// a hardcoded port to make configuration easier.
	app.Listen(":8080")

}
