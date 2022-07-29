package middlewares

import "github.com/gofiber/fiber/v2"

func Auth(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")

	if token != "Bearer token" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"OK":      false,
			"message": "Unauthorized",
		})
	}

	return ctx.Next()
}
