package main

import (
	"github.com/Ntanzi07/gofinance/internal/database"
	"github.com/Ntanzi07/gofinance/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	//connect to database
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	app := fiber.New()

	routes.SetupRoutes(app)
	app.Listen(":8080")

}
