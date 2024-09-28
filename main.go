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

	post_routes := app.Group("/posts")

	post_routes.Post("/", controllers.PostsCreate)
	post_routes.Get("/", controllers.PostsIndex)
	post_routes.Get("/:id", controllers.PostShow)
	post_routes.Put("/:id", controllers.PostUpdate)
	post_routes.Delete("/:id", controllers.PostDelete)

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
