package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Infamous003/ainyx/internal/handler"
)

func Register(app *fiber.App, userHandler *handler.User) {
	app.Get("/health", handler.Health)

	users := app.Group("/users")

	users.Post("/", userHandler.CreateUser)
	users.Get("/", userHandler.ListUsers)
	users.Get("/:id", userHandler.GetUser)
	users.Put("/:id", userHandler.UpdateUser)
	users.Delete("/:id", userHandler.DeleteUser)
}
