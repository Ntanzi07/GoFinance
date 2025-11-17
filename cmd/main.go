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
	//connect to database
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	app := fiber.New()

	routes.SetupRoutes(app, db)
	app.Listen(":8080")

}
