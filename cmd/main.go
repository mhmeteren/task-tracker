package main

import (
	"task-tracker/config"
	"task-tracker/internal/database"
	"task-tracker/internal/di"
	"task-tracker/internal/handler"
	"task-tracker/internal/logger"
	"task-tracker/internal/notifier"
	"task-tracker/internal/router"

	_ "task-tracker/docs"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Task Tracker API
// @version 1.0
// @description This is a sample task logger API
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter token with **Bearer** prefix, e.g. "Bearer eyJhbGciOiJIUzI1..."
// @host 127.0.0.1:3000
// @BasePath /
func main() {
	logger.Init()
	config.LoadConfig()

	logger.GlobalLogger.Info("Application starting...", &logger.LogFields{"tags": []string{"app", "start"}, "environment": config.Cfg.AppEnv, "port": config.Cfg.ServerPort})

	notifier.NotifyInit()

	app := fiber.New(fiber.Config{
		ErrorHandler: handler.FiberErrorHandler,
	})

	database.InitDB()

	if config.Cfg.AppEnv == "development" || config.Cfg.AppEnv == "production" {
		database.AutoMigrateAndSeed()

		app.Get("/swagger/*", fiberSwagger.WrapHandler)
	}

	container := di.InitContainer(database.DB)
	router.AppRoute(app, container)

	app.Listen(":" + config.Cfg.ServerPort)
}
