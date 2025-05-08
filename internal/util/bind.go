package util

import (
	"github.com/gofiber/fiber/v2"
)

func BindAndValidate[T any](c *fiber.Ctx) (*T, error) {
	var req T

	if err := c.BodyParser(&req); err != nil {
		return nil, &BadRequestError{Message: "Request could not reading"}
	}

	if err := ValidateStruct(req); err != nil {
		return nil, err
	}

	return &req, nil
}
