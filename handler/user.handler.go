package handler

import (
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/vorkey/go-fiber-todo/database"
	"github.com/vorkey/go-fiber-todo/model/entity"
	"github.com/vorkey/go-fiber-todo/model/request"
)

func UserHandlerGetAll(ctx *fiber.Ctx) error {
	var users []entity.User
	result := database.DB.Find(&users)

	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(users)
}

func UserHandlerCreate(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)

	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(user)

	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed to store data",
			"error":   strings.Split(errValidate.Error(), "\n"),
		})
	}

	newUser := entity.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
	}

	errCreateUser := database.DB.Create(&newUser).Error

	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newUser,
	})
}
