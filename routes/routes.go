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

	// Rest API 보다 위에 있어야한다. 아래에 :id 가 있어서, id 로 url 을 착각할 수 있다.
	app.Put("/api/user/info", controllers.UpdateInfo)
	app.Put("/api/user/password", controllers.UpdatePassword)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

	// CRUD Rest API Create
	app.Get("/api/users", controllers.AllUsers)
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/users/:id", controllers.GetUser)
	app.Put("/api/users/:id", controllers.UpdateUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)

	// Role CRUD
	app.Get("/api/roles", controllers.AllRoles)
	app.Post("/api/roles", controllers.CreateRole)
	app.Get("/api/roles/:id", controllers.GetRole)
	app.Put("/api/roles/:id", controllers.UpdateRole)
	app.Delete("/api/roles/:id", controllers.DeleteRole)

	app.Get("/api/permissions", controllers.AllPermission)

	// Product CRUD
	app.Get("/api/products", controllers.AllProduct)
	app.Post("/api/products", controllers.CreateProduct)
	app.Get("/api/products/:id", controllers.GetProduct)
	app.Put("/api/products/:id", controllers.UpdateProduct)
	app.Delete("/api/products/:id", controllers.DeleteProduct)

	app.Post("/api/upload", controllers.Upload)
	app.Static("api/uploads", "./uploads")
}
