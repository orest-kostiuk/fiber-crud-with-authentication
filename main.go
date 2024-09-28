package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/orest-kostiuk/fiber-test/config/routes"
	"github.com/orest-kostiuk/fiber-test/database"
)

func main() {
	app := fiber.New()

	database.ConnectToDB()

	routes.SetupRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
