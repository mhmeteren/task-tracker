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

// Login godoc
// @Summary User login
// @Description Authenticates a user and returns access & refresh tokens
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Login credentials"
// @Success 200 {object} dto.AuthResponse
// @Failure 400 {object} util.BadRequestError
// @Failure 422 {object} util.ValidationError
// @Router /api/auth/login [post]
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

// @Summary Refresh Token
// @Description Refresh token, returns access & refresh tokens
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.RefreshTokenRequest true "Refresh credentials"
// @Success 200 {object} dto.AuthResponse
// @Failure 400 {object} util.BadRequestError
// @Failure 401 {object} util.AuthError
// @Failure 422 {object} util.ValidationError
// @Router /api/auth/refresh [post]
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
