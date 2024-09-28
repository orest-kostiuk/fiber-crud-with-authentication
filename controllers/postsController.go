package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/orest-kostiuk/fiber-test/initializers"
	"github.com/orest-kostiuk/fiber-test/models"
)

func PostsCreate(c *fiber.Ctx) error {
	var body struct {
		Title string
		Body  string
	}

	if err := c.BodyParser(&body); err != nil {
		err := c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
		return err
	}

	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		err := c.Status(400).JSON(fiber.Map{"error": "Failed to create post"})
		return err
	}

	err := c.JSON(fiber.Map{"post": post})

	if err != nil {
		return err
	}
	return nil
}

func PostsIndex(c *fiber.Ctx) error {
	var posts []models.Post
	initializers.DB.Find(&posts)

	err := c.JSON(fiber.Map{"posts": posts})

	if err != nil {
		return err
	}
	return nil
}
