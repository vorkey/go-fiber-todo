package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vorkey/go-fiber-todo/route"
	// "go-fiber-todo/handlers"
)

func main() {
	app := fiber.New()

	route.RouteInit(app)

	app.Listen(":3000")
}
