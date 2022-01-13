package main

import "github.com/gofiber/fiber/v2"

func handler(c *fiber.Ctx) error {
	return c.SendString("hello from another func")
}

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! again!")
	})
	app.Get("/s", handler)
	app.Listen(":3000")
}
