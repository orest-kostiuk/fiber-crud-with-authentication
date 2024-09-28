package sessionsRoutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/orest-kostiuk/fiber-test/app/controllers/sessionsController"
)

func SetupSessionsRoutes(router fiber.Router) {
	router.Post("/login", sessionsController.Login)
}
