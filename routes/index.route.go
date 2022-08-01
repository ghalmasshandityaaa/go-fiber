package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go-fiber-api/config"
	"go-fiber-api/handlers"
	"go-fiber-api/middlewares"
	"go-fiber-api/utils"
)

func RouteInitialization(route *fiber.App) {
	/** Static route */
	route.Static("/public", config.ProjectRootPath+"./public/assets")

	/** versioning api */
	api := route.Group("/api/v1", logger.New())

	/** Grouping users route */
	users := api.Group("/users", logger.New())
	auth := api.Group("/auth", logger.New())
	book := api.Group("/books", logger.New())

	/** Users route */
	users.Get("/", middlewares.Auth, handlers.AllUsers)
	users.Post("/", handlers.CreateUser)
	users.Get("/:id", handlers.GetUserById)
	users.Put("/:id", handlers.UpdateUserById)
	users.Delete("/:id", handlers.DeleteUserById)

	/** Auth route */
	auth.Post("/login", handlers.Login)

	/** Book route */
	book.Post("/", utils.UploadFile, handlers.CreateBook)
}
