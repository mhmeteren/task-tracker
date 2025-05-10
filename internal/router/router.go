package router

import (
	"task-tracker/internal/di"
	"task-tracker/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func AppRoute(app *fiber.App, container *di.Container) {
	api := app.Group("/api", middleware.RateLimitByIP())
	RegisterUserRoutes(api, container.UserController)
	RegisterAuthRoutes(api, container.AuthController)
	RegisterTaskRoutes(api, container.TaskController)
	RegisterLogRoutes(api, container.LogController)
	RegisterTaskNotificationRoutes(api, container.TaskNotificationController)
}
