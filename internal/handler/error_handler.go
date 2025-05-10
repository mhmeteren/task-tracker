package handler

import (
	"task-tracker/internal/util"

	"github.com/gofiber/fiber/v2"
)

func FiberErrorHandler(c *fiber.Ctx, err error) error {
	statusCode := getStatusCode(err)
	errorBody := getErrorBody(err)

	return c.Status(statusCode).JSON(errorBody)
}

func getErrorBody(err error) *util.ErrorResponse[any] {
	switch e := err.(type) {

	case *util.ValidationError:
		return &util.ErrorResponse[any]{Error: e.Errors}

	case *util.NotFoundError, *util.ConflictError, *util.AuthError, *util.BadRequestError, *util.RateLimitError:
		return &util.ErrorResponse[any]{Error: e.Error()}

	case *fiber.Error:
		return &util.ErrorResponse[any]{Error: e.Message}

	default:
		return &util.ErrorResponse[any]{Error: "Server error. message: " + e.Error()}
	}
}

func getStatusCode(err error) int {
	switch e := err.(type) {
	case *util.ValidationError:
		return fiber.StatusUnprocessableEntity
	case *util.NotFoundError:
		return fiber.StatusNotFound
	case *util.ConflictError:
		return fiber.StatusConflict
	case *util.AuthError:
		return fiber.StatusUnauthorized
	case *util.BadRequestError:
		return fiber.StatusBadRequest
	case *util.RateLimitError:
		return fiber.StatusTooManyRequests
	case *fiber.Error:
		return e.Code
	default:
		return fiber.StatusInternalServerError
	}
}
