package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/orest-kostiuk/fiber-test/app/models"
	"github.com/orest-kostiuk/fiber-test/database"
)

func PostsCreate(c *fiber.Ctx) error {
	db := database.DB

	var body struct {
		Title string
		Body  string
	}

	if err := c.BodyParser(&body); err != nil {
		err := c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
		return err
	}

	post := models.Post{Title: body.Title, Body: body.Body}
	result := db.Create(&post)

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
	db := database.DB

	var posts []models.Post
	db.Find(&posts)

	err := c.JSON(fiber.Map{"posts": posts})

	if err != nil {
		return err
	}
	return nil
}

func PostShow(c *fiber.Ctx) error {
	db := database.DB

	id := c.Params("id")
	var post models.Post
	db.Find(&post, id)

	err := c.JSON(fiber.Map{"post": post})

	if err != nil {
		return err
	}
	return nil
}

func PostUpdate(c *fiber.Ctx) error {
	db := database.DB

	id := c.Params("id")

	var body struct {
		Title string
		Body  string
	}

	if err := c.BodyParser(&body); err != nil {
		err := c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
		return err
	}

	var post models.Post
	db.Find(&post, id)

	post.Title = body.Title
	post.Body = body.Body

	result := db.Save(&post)

	if result.Error != nil {
		err := c.Status(400).JSON(fiber.Map{"error": "Failed to update post"})
		return err
	}

	err := c.JSON(fiber.Map{"post": post})

	if err != nil {
		return err
	}
	return nil
}

func PostDelete(c *fiber.Ctx) error {
	db := database.DB

	id := c.Params("id")
	var post models.Post
	db.Find(&post, id)

	result := db.Delete(&post)

	if result.Error != nil {
		err := c.Status(400).JSON(fiber.Map{"error": "Failed to delete post"})
		return err
	}

	err := c.JSON(fiber.Map{"message": "Post deleted"})

	if err != nil {
		return err
	}
	return nil
}
