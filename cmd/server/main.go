package main

import (
	"github.com/Infamous003/ainyx/config"
	"github.com/Infamous003/ainyx/internal/database"
	"github.com/Infamous003/ainyx/internal/logger"
	"github.com/Infamous003/ainyx/internal/routes"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {
	logr := logger.New()
	defer logr.Sync()

	cfg, err := config.Load()
	if err != nil {
		logr.Fatal("failed to load config", zap.Error(err))
	}

	db, err := database.Connect(cfg)
	if err != nil {
		logr.Fatal("failed to connect to DB", zap.Error(err))
	}
	defer db.Close()

	app := fiber.New()

	routes.Register(app)

	logr.Info("server running", zap.String("port", cfg.Port))
	if err := app.Listen(":" + cfg.Port); err != nil {
		logr.Fatal("failed to start the server", zap.Error(err))
	}
}
