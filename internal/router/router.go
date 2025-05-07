package router

import (
	"task-tracker/internal/di"

	"github.com/gofiber/fiber/v2"
)

func AppRoute(app *fiber.App, container *di.Container) {
	api := app.Group("/api")
	RegisterUserRoutes(api, container.UserController)
	RegisterAuthRoutes(api, container.AuthController)
	RegisterTaskRoutes(api, container.TaskController)
	RegisterLogRoutes(api, container.LogController)
}
