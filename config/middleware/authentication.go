package middleware

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/orest-kostiuk/fiber-test/app/models"
	"github.com/orest-kostiuk/fiber-test/database"
	"net/http"
	"os"
	"time"
)

func RequireAuth(c *fiber.Ctx) error {
	db := database.DB

	tokenString := c.Cookies("Authorization")

	if tokenString == "" {
		err := c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		return err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			err := c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
			return err
		} else {
			err := c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Bad request"})
			return err
		}
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			err := c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Token expired"})
			return err
		}

		var user models.User
		db.Find(&user, claims["sub"])

		if user.ID == uuid.Nil {
			err := c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
			return err
		}

		c.Locals("user", user)

		err := c.Next()
		if err != nil {
			return err
		}
	} else {
		err := c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		return err
	}
	return nil
}
