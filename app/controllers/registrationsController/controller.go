package registrationsController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/orest-kostiuk/fiber-test/app/models"
	"github.com/orest-kostiuk/fiber-test/database"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Signup(c *fiber.Ctx) error {
	db := database.DB

	var body struct {
		Email    string
		Password string
	}

	if err := c.BodyParser(&body); err != nil {
		err := c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		err := c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Internal server error"})
		return err
	}

	user := models.User{Email: body.Email, Password: string(hash)}
	result := db.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return result.Error
	}

	err = c.Status(http.StatusOK).JSON(fiber.Map{
		"user": user,
	})

	if err != nil {
		return err
	}
	return nil
}
