package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/handlers"
)

func RouteInitialization(route *fiber.App) {
	route.Get("/users", handlers.AllUsers)
}
