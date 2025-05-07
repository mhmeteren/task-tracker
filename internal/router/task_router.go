package router

import (
	"task-tracker/internal/controller"
	"task-tracker/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterTaskRoutes(router fiber.Router, controller *controller.TaskController) {
	taskGroup := router.Group("/tasks")
	taskGroup.Post("/", middleware.Authorize("user"), controller.CreateNewTask)
	taskGroup.Put("/:id", middleware.Authorize("user"), controller.UpdateTask)
	taskGroup.Delete("/:id", middleware.Authorize("user"), controller.DeleteTask)
	taskGroup.Get("/", middleware.Authorize("user"), controller.GetAllByLoggedUser)

}
