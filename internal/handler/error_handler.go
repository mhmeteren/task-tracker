package handler

import (
	"task-tracker/internal/context"
	"task-tracker/internal/logger"
	"task-tracker/internal/util"

	"github.com/gofiber/fiber/v2"
)

func FiberErrorHandler(c *fiber.Ctx, err error) error {
	statusCode := getStatusCode(err)
	errorBody := getErrorBody(err)

	addLog(c, err, statusCode)

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

func addLog(c *fiber.Ctx, err error, statusCode int) {
	var user_id *uint
	userCtx := context.GetUserContext(c)
	if userCtx != nil {
		user_id = &userCtx.UserID
	}

	logger.GlobalLogger.Error("Request failed", &logger.LogFields{
		"path":       c.Path(),
		"method":     c.Method(),
		"status":     statusCode,
		"error":      err.Error(),
		"client_ip":  c.IP(),
		"user_id":    user_id,
		"user_agent": c.Get("User-Agent"),
		"tags":       []string{"request", "handler"},
	})
}
