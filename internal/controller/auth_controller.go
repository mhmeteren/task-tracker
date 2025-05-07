package controller

import (
	"task-tracker/internal/dto"
	"task-tracker/internal/service"
	"task-tracker/internal/util"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	AuthService service.AuthService
}

func NewAuthController(service service.AuthService) *AuthController {
	return &AuthController{AuthService: service}
}

func (ac *AuthController) Login(c *fiber.Ctx) error {
	req, err := util.BindAndValidate[dto.LoginRequest](c)
	if err != nil {
		return err
	}

	result, err := ac.AuthService.Login(req)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	return c.JSON(result)
}

func (ac *AuthController) RefreshToken(c *fiber.Ctx) error {
	req, err := util.BindAndValidate[dto.RefreshTokenRequest](c)
	if err != nil {
		return err
	}

	result, err := ac.AuthService.RefreshToken(req)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	return c.JSON(result)
}
