package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/vorkey/go-fiber-todo/database"
	"github.com/vorkey/go-fiber-todo/model/entity"
)

func UserHandlerGetAll(ctx *fiber.Ctx) error {
	var users []entity.User
	result := database.DB.Find(&users)

	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(users)
}
