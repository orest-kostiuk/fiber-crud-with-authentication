package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/orest-kostiuk/fiber-test/config/routes/postRoutes"
)

func SetupRoutes(app *fiber.App) {
	application := app.Group("", logger.New())

	postRoutes.SetupPostRoutes(application)
}
