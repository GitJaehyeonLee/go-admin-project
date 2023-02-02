package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-admin/controllers"
	"go-admin/middlewares"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	// 밑에는 토큰 검증 필요 / 위는 토큰 검증 불필요.
	app.Use(middlewares.IsAuthenticated)

	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

	// CRUD Rest API Create
	app.Get("/api/users", controllers.AllUsers)
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/users/:id", controllers.GetUser)
	app.Put("/api/users/:id", controllers.UpdateUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)
}
