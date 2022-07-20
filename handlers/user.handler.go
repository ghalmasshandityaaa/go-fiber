package handlers

import (
	"fmt"
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

func CreateUser(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)

	/** check body */
	if err := ctx.BodyParser(user); err != nil {
		fmt.Println("error => ", err)
		return ctx.Status(400).JSON(fiber.Map{
			"OK":      false,
			"message": err.Error(),
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
		return ctx.Status(400).JSON(fiber.Map{
			"OK":      false,
			"message": errInsert,
			"data":    nil,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"OK":      true,
		"message": "Success create user",
		"data":    user,
	})
}
