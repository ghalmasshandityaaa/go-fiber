package handlers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/database"
	"go-fiber-api/models/entity"
	"go-fiber-api/models/request"
	"strings"
)

func Login(ctx *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)

	/** check body */
	if err := ctx.BodyParser(loginRequest); err != nil {
		fmt.Println("error => ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"OK":      false,
			"message": err.Error(),
		})
	}

	/** Validate request body */
	validate := validator.New()
	errValidate := validate.Struct(loginRequest)

	if errValidate != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"OK":      false,
			"message": strings.Split(errValidate.Error(), "\n"),
		})
	}

	var user entity.User
	err := database.DB.Debug().First(&user, "email = ?", loginRequest.Email).Error
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"OK":      false,
			"message": "Sorry your email is not registered!",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"OK":      true,
		"message": "token",
	})
}
