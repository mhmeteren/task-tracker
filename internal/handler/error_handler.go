package handler

import (
	"task-tracker/internal/util"

	"github.com/gofiber/fiber/v2"
)

func FiberErrorHandler(c *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case *util.ValidationError:
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"validation_errors": e.Errors,
		})
	case *util.NotFoundError:
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": e.Error(),
		})
	case *util.ConflictError:
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": e.Error(),
		})
	case *util.AuthError:
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": e.Error(),
		})
	case *util.BadRequestError:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": e.Error(),
		})
	case *fiber.Error:
		return c.Status(e.Code).JSON(fiber.Map{"error": e.Message})
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Server error. message: " + e.Error(),
		})
	}
}
