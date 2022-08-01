package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/utils"
)

func Auth(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"OK":      false,
			"message": "Invalid token!",
		})
	}

	//_, err := utils.VerifyAccessToken(token)
	claims, err := utils.DecodeToken(token)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"OK":      false,
			"message": "Unauthorized",
		})
	}

	if claims["role"].(string) != "admin" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"OK":      false,
			"message": "Forbidden access",
		})
	}

	ctx.Locals("auth", claims)

	return ctx.Next()
}
