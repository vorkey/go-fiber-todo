package main

import (
	"github.com/gofiber/fiber/v2"
	// "go-fiber-todo/handlers"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"Message": "Welcome to Go API!",
		})

	})

	app.Listen(":3000")
}
