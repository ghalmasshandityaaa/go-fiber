package handlers

import "github.com/gofiber/fiber/v2"

func AllUsers(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{
		"OK":      true,
		"message": "Success get all users data",
		"data":    nil,
	})
}
