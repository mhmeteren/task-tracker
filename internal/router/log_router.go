package router

import (
	"task-tracker/internal/controller"
	"task-tracker/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterLogRoutes(router fiber.Router, controller *controller.LogController) {
	logGroup := router.Group("/logs")
	logGroup.Get("/:taskKey/:taskSecret", middleware.RateLimitByTask(), controller.AddLog)

	logGroup.Get("/:taskID", middleware.Authorize("user"), controller.GetAllByTask)
}
