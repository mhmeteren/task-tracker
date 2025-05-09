package router

import (
	"task-tracker/internal/controller"
	"task-tracker/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterTaskNotificationRoutes(router fiber.Router, controller *controller.TaskNotificationController) {
	taskNotificationGroup := router.Group("/task-notifications")
	taskNotificationGroup.Post("/", middleware.Authorize("user"), controller.NotificationServiceIntegration)
	taskNotificationGroup.Delete("/:taskID", middleware.Authorize("user"), controller.DeleteNotificationServiceInformation)
}
