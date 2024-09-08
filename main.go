package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vorkey/go-fiber-todo/database"
	"github.com/vorkey/go-fiber-todo/database/migration"
	"github.com/vorkey/go-fiber-todo/route"
)

func main() {
	database.GetMysqlInstance()
	migration.RunMigration()

	app := fiber.New()

	route.RouteInit(app)

	app.Listen(":3000")
}
