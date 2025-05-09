package main

import (
	"task-tracker/config"
	"task-tracker/internal/database"
	"task-tracker/internal/di"
	"task-tracker/internal/handler"
	"task-tracker/internal/notifier"
	"task-tracker/internal/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadConfig()
	notifier.NotifyInit()

	app := fiber.New(fiber.Config{
		ErrorHandler: handler.FiberErrorHandler,
	})

	database.InitDB()
	if config.Cfg.AppEnv == "development" {
		database.AutoMigrateAndSeed()
	}

	container := di.InitContainer(database.DB)
	router.AppRoute(app, container)

	app.Listen(":" + config.Cfg.ServerPort)
}
