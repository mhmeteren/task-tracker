package main

import (
	"task-tracker/config"
	"task-tracker/internal/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadConfig()

	app := fiber.New(fiber.Config{})

	database.InitDB()
	if config.Cfg.AppEnv == "development" {
		database.AutoMigrateAndSeed()
	}

	app.Listen(":" + config.Cfg.ServerPort)
}
