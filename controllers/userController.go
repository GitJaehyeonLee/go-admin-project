package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-admin/database"
	"go-admin/models"
	"golang.org/x/crypto/bcrypt"
)

// User 관련 기능 정의

func AllUsers(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Find(&users)

	return c.JSON(&users)
}

// 관리자가 사용자 만드는 기능
func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte("1234"), 14)

	user.Password = password

	database.DB.Create(&user)

	return c.JSON(user)
}
