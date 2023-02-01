package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-admin/models"
)

func Register(c *fiber.Ctx) error {
	user := models.User{
		FirstName: "John",
	}

	user.LastName = "Doe"
	user.Email = "example.com"

	return c.JSON(user)
}
