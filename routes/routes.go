package routes

import (
	"github.com/gofiber/fiber/v2"
	"go_admin.com/controllers"
	middleware "go_admin.com/middlewares"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middleware.IsAuthenticated)

	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
	app.Get("/api/users", controllers.AllUsers)
	app.Post("/api/create-user", controllers.CreateUser)
}
