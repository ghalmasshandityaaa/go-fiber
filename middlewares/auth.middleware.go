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

	_, err := utils.VerifyAccessToken(token)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"OK":      false,
			"message": "Unauthorized",
		})
	}

	return ctx.Next()
}
