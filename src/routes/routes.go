package routes

import "github.com/gofiber/fiber/v2"

func Main() fiber.Handler {
	greet := "Welcome to sapangpera.online. Site is under construction. "
	return func(c *fiber.Ctx) error {
		return c.SendString(greet)
	}
}

func Api(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{"message": "hello", "status": "OK"})
}
