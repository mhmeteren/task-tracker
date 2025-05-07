package router

import (
	"task-tracker/internal/controller"

	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(router fiber.Router, authController *controller.AuthController) {
	auth := router.Group("/auth")
	auth.Post("/login", authController.Login)
	auth.Post("/refresh", authController.RefreshToken)
}
