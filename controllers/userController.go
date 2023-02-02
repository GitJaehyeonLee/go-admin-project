package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-admin/database"
	"go-admin/models"
	"strconv"
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

	user.SetPassword("1234")

	database.DB.Create(&user)

	return c.JSON(user)
}

func GetUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	database.DB.Find(&user)

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user := models.User{
		Id: uint(id),
	}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	database.DB.Model(&user).Updates(user)

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user := models.User{
		Id: uint(id),
	}

	database.DB.Delete(&user)

	return nil
}
