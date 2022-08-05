package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go-fiber-api/database"
	"go-fiber-api/database/migration"
	"go-fiber-api/routes"
	"log"
	"os"
)

func main() {
	/** Database Initialization */
	database.DBInitialization()

	/** Run Migration */
	migration.RunMigration()

	app := fiber.New(fiber.Config{
		AppName: "go-fiber-api",
	})

	app.Use(cors.New(cors.ConfigDefault))

	/** Route Initialization */
	routes.RouteInitialization(app)

	/** 404 Error Handler */
	app.Use(func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"OK":      false,
			"message": "Ooppss! 404 api not found!",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Fatal(app.Listen(":" + port))
}
