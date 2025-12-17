package main

import (
	"log"

	"github.com/Infamous003/ainyx/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.Register(app)

	if err := app.Listen(":8080"); err != nil {
		log.Fatal("failed to start the server")
	}
}
