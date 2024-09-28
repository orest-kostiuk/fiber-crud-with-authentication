package registraionRoutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/orest-kostiuk/fiber-test/app/controllers/registrationsController"
)

func SetupRegistrationRoutes(router fiber.Router) {
	router.Post("/signup", registrationsController.Signup)
}
