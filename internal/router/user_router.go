package router

import (
	"task-tracker/internal/controller"
	"task-tracker/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(router fiber.Router, controller *controller.UserController) {
	userGroup := router.Group("/users")

	userGroup.Post("/", controller.CreateUser)
	userGroup.Get("/profile", middleware.Authorize(), controller.GetProfile)
	userGroup.Get("/", middleware.Authorize("admin"), controller.GetAllUsers)

	userGroup.Delete("/:id", middleware.Authorize("admin"), controller.DeleteUser)
}
