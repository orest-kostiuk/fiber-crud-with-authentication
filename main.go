package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/orest-kostiuk/fiber-test/controllers"
	"github.com/orest-kostiuk/fiber-test/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	app := fiber.New()

	app.Post("/posts", controllers.PostsCreate)
	app.Get("/posts", controllers.PostsIndex)
	app.Get("/posts/:id", controllers.PostShow)
	app.Put("/posts/:id", controllers.PostUpdate)

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
