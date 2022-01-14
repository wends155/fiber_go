package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/wends155/fiber_go/src/routes"
)

func New() *fiber.App {
	app := fiber.New()
	app.Use(compress.New())

	app.Static("/", "./frontend/build")

	app.Get("/s", routes.Api)
	return app
}
