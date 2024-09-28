package sessionsController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/orest-kostiuk/fiber-test/app/models"
	"github.com/orest-kostiuk/fiber-test/database"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func Login(c *fiber.Ctx) error {
	db := database.DB

	var body struct {
		Email    string
		Password string
	}

	if err := c.BodyParser(&body); err != nil {
		err := c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		return err
	}

	var user models.User
	dbErr := db.First(&user, "email = ?", body.Email)

	if dbErr.Error != nil {
		err := c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		return err
	}

	if user.ID == uuid.Nil {
		err := c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
		return err
	}

	bcryptErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if bcryptErr != nil {
		err := c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
		return err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, signError := token.SignedString([]byte(os.Getenv("SECRET")))

	if signError != nil {
		err := c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Fail to create token"})
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:     "Authorization",
		Value:    tokenString,
		Expires:  time.Now().Add(3600 * 24 * 30 * time.Second),
		Path:     "/",
		SameSite: "Lax",
		Secure:   false,
		HTTPOnly: true,
	})

	err := c.Status(http.StatusOK).JSON(fiber.Map{
		"token": tokenString,
	})
	if err != nil {
		return err
	}
	return nil
}
