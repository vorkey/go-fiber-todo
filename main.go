package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vorkey/go-fiber-todo/database"
	"github.com/vorkey/go-fiber-todo/route"
)

func main() {
	database.GetMysqlInstance()

	app := fiber.New()

	route.RouteInit(app)

	app.Listen(":3000")
}
