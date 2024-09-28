package postRoutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/orest-kostiuk/fiber-test/app/controllers/postsController"
)

func SetupPostRoutes(router fiber.Router, auth func(c *fiber.Ctx) error) {
	postRoutes := router.Group("/posts", auth)

	postRoutes.Post("/", postsController.PostsCreate)
	postRoutes.Get("/", postsController.PostsIndex)
	postRoutes.Get("/:id", postsController.PostShow)
	postRoutes.Put("/:id", postsController.PostUpdate)
	postRoutes.Delete("/:id", postsController.PostDelete)
}
