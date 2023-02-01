package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go-admin/database"
	"go-admin/routes"
)

func main() {
	database.Connect()
	app := fiber.New()
	// cookie 데이터 접근 허용
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":3000")
}

// air 모듈 추가함 -> 실행 air
