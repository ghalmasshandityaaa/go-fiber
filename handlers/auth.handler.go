package handlers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/database"
	"go-fiber-api/models/entity"
	"go-fiber-api/models/request"
	"go-fiber-api/utils"
	"time"
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

	/** Check Email in database */
	var user entity.User
	err := database.DB.Debug().First(&user, "email = ?", loginRequest.Email).Error
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"OK":      false,
			"message": "Sorry your email is not registered!",
		})
	}

	/** Validate email and password */
	isValid := utils.VerifyPassword(user.Password, loginRequest.Password)
	if !isValid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"OK":      false,
			"message": "Wrong credentials",
		})
	}

	/** Generate JWT Token */
	claims := jwt.MapClaims{
		"name":    user.Name,
		"email":   user.Email,
		"address": user.Address,
		"exp":     time.Now().Add(time.Minute * 5).Unix(),
		"role":    "user",
	}

	if user.Email == "ghalmas@gmail.com" {
		claims["role"] = "admin"
	}

	token, errGenerateToken := utils.GenerateAccessToken(&claims)
	if errGenerateToken != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"OK":      false,
			"message": "Internal server error!",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"OK":      true,
		"message": token,
	})
}
