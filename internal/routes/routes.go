package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Infamous003/ainyx/internal/handler"
)

func Register(app *fiber.App) {
	app.Get("/health", handler.Health)
}
