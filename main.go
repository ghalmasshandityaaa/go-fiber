package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "go-fiber-api",
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"OK":      true,
			"message": "Server is running",
			"data":    nil,
		})
	})

	app.Listen(":6000")
}
