package util

import (
	"task-tracker/internal/parameter"

	"github.com/gofiber/fiber/v2"
)

func BindAndSetDefaultParameters[T any, PT interface {
	*T
	parameter.Parameters
}](c *fiber.Ctx) (PT, error) {
	var params T
	ptr := PT(&params)

	if err := c.QueryParser(ptr); err != nil {
		return nil, &BadRequestError{Message: "Invalid query params"}
	}

	ptr.SetDefaults()
	return ptr, nil
}
