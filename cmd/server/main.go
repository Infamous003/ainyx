package main

import (
	"log"

	"github.com/Infamous003/ainyx/config"
	"github.com/Infamous003/ainyx/internal/database"
	"github.com/Infamous003/ainyx/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("failed to load config: ", err.Error())
	}

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("failed to connect to PostgreSQL: ", err.Error())
	}
	defer db.Close()
	log.Println("successfully connected to PostgreSQL")

	app := fiber.New()

	routes.Register(app)

	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatal("failed to start the server")
	}
}
