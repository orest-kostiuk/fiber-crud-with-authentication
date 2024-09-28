package postRoutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/orest-kostiuk/fiber-test/app/controllers/postsController"
)

func SetupPostRoutes(router fiber.Router) {
	postRoutes := router.Group("/posts")

	postRoutes.Post("/", postsController.PostsCreate)
	postRoutes.Get("/", postsController.PostsIndex)
	postRoutes.Get("/:id", postsController.PostShow)
	postRoutes.Put("/:id", postsController.PostUpdate)
	postRoutes.Delete("/:id", postsController.PostDelete)
}
