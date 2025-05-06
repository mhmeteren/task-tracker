package router

import (
	"task-tracker/internal/controller"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(router fiber.Router, controller *controller.UserController) {
	router.Post("/users", controller.CreateUser)
}
