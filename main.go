package main

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/routes"
	"log"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "go-fiber-api",
	})

	/** Route Initialization */
	routes.RouteInitialization(app)

	log.Fatal(app.Listen(":6000"))
}
