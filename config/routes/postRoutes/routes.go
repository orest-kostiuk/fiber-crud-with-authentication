package postRoutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/orest-kostiuk/fiber-test/app/controllers"
)

func SetupPostRoutes(router fiber.Router) {
	postRoutes := router.Group("/posts")

	postRoutes.Post("/", controllers.PostsCreate)
	postRoutes.Get("/", controllers.PostsIndex)
	postRoutes.Get("/:id", controllers.PostShow)
	postRoutes.Put("/:id", controllers.PostUpdate)
	postRoutes.Delete("/:id", controllers.PostDelete)
}
