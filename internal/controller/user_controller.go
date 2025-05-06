package controller

import (
	"task-tracker/internal/dto"
	"task-tracker/internal/service"
	"task-tracker/internal/util"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	service service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{service}
}

func (ctl *UserController) CreateUser(c *fiber.Ctx) error {
	req, err := util.BindAndValidate[dto.CreateUserRequest](c)
	if err != nil {
		return err
	}

	checked_user, _ := ctl.service.GetByEmail(req.Email)
	if checked_user != nil {
		return fiber.NewError(fiber.StatusNotFound, "Email must be unique")
	}

	user := req.ToModel()

	if err := ctl.service.Create(&user, req.Password); err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ResultResponse{Message: "User successfully created"})
}
