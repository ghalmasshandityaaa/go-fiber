package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go-fiber-api/handlers"
)

func RouteInitialization(route *fiber.App) {
	/** versioning api */
	api := route.Group("/api/v1", logger.New())

	/** Grouping users route */
	users := api.Group("/users", logger.New())

	users.Get("/", handlers.AllUsers)
	users.Post("/", handlers.CreateUser)
	users.Get("/:id", handlers.GetUserById)
	users.Put("/:id", handlers.UpdateUserById)
}
