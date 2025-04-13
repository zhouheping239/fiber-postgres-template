package main

import (
	"backend_project/api"
	config "backend_project/configs"
	"backend_project/internal/middleware/exception"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	config.Init()

	app := fiber.New(fiber.Config{
		AppName:                 "backend_project",
		BodyLimit:               4 * 1024 * 1024,
		EnableTrustedProxyCheck: true,
		ErrorHandler:            exception.ApiException,
		Prefork:                 true,
		StrictRouting:           true,
	})

	app.Use(cors.New())

	api.SetupRoutes(app)

	app.Listen(":8080")
}
