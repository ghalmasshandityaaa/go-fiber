package handlers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/database"
	"go-fiber-api/models/entity"
	"go-fiber-api/models/request"
	"go-fiber-api/models/response"
	"go-fiber-api/utils"
	"log"
)

func AllUsers(ctx *fiber.Ctx) error {
	var users []entity.User

	result := database.DB.Debug().Find(&users)

	if result.Error != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"OK":      false,
			"message": result.Error,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
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
		})
	}

	/** Check email is exist or not */

	/** Insert user into database */
	newUser := entity.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
		Age:     user.Age,
	}

	hashedPassword, errHash := utils.HashingPassword(user.Password)
	if errHash != nil {
		log.Println("Hashing Password Error => ", errHash)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"OK":      false,
			"message": "Internal server error!",
		})
	}

	newUser.Password = hashedPassword

	errInsert := database.DB.Create(&newUser).Error
	if errInsert != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"OK":      false,
			"message": errInsert,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"OK":      true,
		"message": "Success create user",
		"data":    user,
	})
}

func GetUserById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	var users entity.User

	err := database.DB.First(&users, "id = ?", userId).Error

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"OK":      false,
			"message": err,
		})
	}

	userResponse := response.UserResponse{
		ID:        users.ID,
		Name:      users.Name,
		Email:     users.Email,
		Age:       users.Age,
		Address:   users.Address,
		Phone:     users.Phone,
		CreatedAt: users.CreatedAt,
		UpdatedAt: users.CreatedAt,
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"OK":               true,
		"message":          "Success get users data",
		"dataWithResponse": userResponse,
		"dataUsers":        users,
	})
}

func UpdateUserById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	var users entity.User
	user := new(request.UserUpdateRequest)

	err := database.DB.Debug().First(&users, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"OK":      false,
			"message": "User is not exist",
		})
	}

	/** check body */
	if errBody := ctx.BodyParser(user); errBody != nil {
		fmt.Println("error => ", errBody)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"OK":      false,
			"message": errBody.Error(),
		})
	}

	/** Validate request body */
	validate := validator.New()
	errValidate := validate.Struct(user)

	if errValidate != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"OK":      false,
			"message": errValidate.Error(),
		})
	}

	/** Update user data */
	if user.Name != "" {
		users.Name = user.Name
	}
	users.Age = user.Age
	users.Address = user.Address
	users.Phone = user.Phone

	errUpdate := database.DB.Debug().Save(&users).Error
	if errUpdate != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"OK":      false,
			"message": errUpdate,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"OK":      true,
		"message": "Success update users data",
	})
}

func DeleteUserById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	var users entity.User

	err := database.DB.First(&users, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"OK":      false,
			"message": "Sorry user not found",
		})
	}

	errDelete := database.DB.Debug().Delete(&users).Error
	if errDelete != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"OK":      false,
			"message": errDelete,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"OK":      true,
		"message": "Success deleted user",
	})
}
