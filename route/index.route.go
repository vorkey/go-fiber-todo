package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vorkey/go-fiber-todo/handler"
)

func RouteInit(r *fiber.App) {

	r.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"Message": "Welcome to Go API!",
		})
	})

	r.Get("/user", handler.UserHandlerGetAll)
	r.Post("/user", handler.UserHandlerCreate)
}
