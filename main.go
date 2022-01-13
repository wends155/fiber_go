package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

func handler(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{"message": "hello", "status": "OK"})
}

func main() {
	app := fiber.New()
	app.Use(compress.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! again!")
	})
	app.Get("/s", handler)
	app.Listen(":3000")
}
