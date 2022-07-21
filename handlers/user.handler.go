package handlers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/database"
	"go-fiber-api/models/entity"
	"go-fiber-api/models/request"
	"log"
)

func AllUsers(ctx *fiber.Ctx) error {
	var users []entity.User

	result := database.DB.Debug().Find(&users)

	if result.Error != nil {
		log.Println(result.Error)
	}

	return ctx.Status(200).JSON(fiber.Map{
		"OK":      true,
		"message": "Success get all users data",
		"data":    users,
	})
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func CreateUser(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)

	/** check body */
	if err := ctx.BodyParser(user); err != nil {
		fmt.Println("error => ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"OK":      false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	/** Validate request body */
	validate := validator.New()
	errValidate := validate.Struct(user)

	if errValidate != nil {
		var errors []*ErrorResponse
		for _, err := range errValidate.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"OK":      false,
			"message": errors,
			"data":    nil,
		})
	}

	/** Insert user into database */
	newUser := entity.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
		Age:     user.Age,
	}

	errInsert := database.DB.Create(&newUser).Error
	if errInsert != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"OK":      false,
			"message": errInsert,
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"OK":      true,
		"message": "Success create user",
		"data":    user,
	})
}
