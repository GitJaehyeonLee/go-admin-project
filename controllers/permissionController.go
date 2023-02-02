package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-admin/database"
	"go-admin/models"
)

func AllPermission(c *fiber.Ctx) error {
	var Permissions []models.Permission

	database.DB.Find(&Permissions)

	return c.JSON(Permissions)
}
