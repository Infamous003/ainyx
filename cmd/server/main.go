package main

import (
	"log"

	"github.com/Infamous003/ainyx/config"
	"github.com/Infamous003/ainyx/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("failed to load config: ", err.Error())
	}

	app := fiber.New()

	routes.Register(app)

	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatal("failed to start the server")
	}
}
