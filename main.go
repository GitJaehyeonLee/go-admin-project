package main

import (
	"github.com/gofiber/fiber/v2"
	"go-admin/database"
	"go-admin/routes"
)

func main() {
	database.Connect()
	app := fiber.New()

	routes.Setup(app)

	app.Listen(":3000")
}

// air 모듈 추가함 -> 실행 air
