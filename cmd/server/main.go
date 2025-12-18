package main

import (
	"github.com/Infamous003/ainyx/config"
	"github.com/Infamous003/ainyx/internal/database"
	"github.com/Infamous003/ainyx/internal/database/sqlc"
	"github.com/Infamous003/ainyx/internal/handler"
	"github.com/Infamous003/ainyx/internal/logger"
	"github.com/Infamous003/ainyx/internal/repository"
	"github.com/Infamous003/ainyx/internal/routes"
	"github.com/Infamous003/ainyx/internal/service"
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

	querier := sqlc.New(db)
	userRepo := repository.NewUser(querier)
	userService := service.NewUser(userRepo)
	userHandler := handler.NewUser(userService, logr)

	app := fiber.New()

	routes.Register(app, userHandler)

	logr.Info("server running", zap.String("port", cfg.Port))
	if err := app.Listen(":" + cfg.Port); err != nil {
		logr.Fatal("failed to start the server", zap.Error(err))
	}
}
