package router

import (
	"task-tracker/internal/controller"

	"github.com/gofiber/fiber/v2"
)

func RegisterHealthCheckRoutes(router fiber.Router, healthCheckController *controller.HealthCheckController) {
	router.Get("/health", healthCheckController.Health)
}
