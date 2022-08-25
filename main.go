package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go_admin.com/database"
	"go_admin.com/routes"
)

func main() {

	database.Connect()

	app := fiber.New()

	// Important for the front end to receive a cookie
	// created during login
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":3000")
}
