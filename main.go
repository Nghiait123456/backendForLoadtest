package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("homepage")
	})

	app.Listen(":8000")
}
