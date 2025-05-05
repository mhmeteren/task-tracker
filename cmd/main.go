package main

import (
	"task-tracker/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadConfig()

	app := fiber.New(fiber.Config{})

	app.Listen(":" + config.Cfg.ServerPort)
}
