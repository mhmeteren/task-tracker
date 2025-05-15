package controller

import (
	"strconv"
	"task-tracker/internal/context"
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
		return fiber.NewError(fiber.StatusBadRequest, "Email must be unique")
	}

	user := req.ToModel()

	if err := ctl.service.Create(&user, req.Password); err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ResultResponse{Message: "User successfully created"})
}

func (ctl *UserController) GetProfile(c *fiber.Ctx) error {
	userCtx := context.GetUserContext(c)

	user, err := ctl.service.GetProfile(userCtx.UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(dto.ToUserDetail(*user))
}

func (ctl *UserController) GetAllUsers(c *fiber.Ctx) error {
	users, err := ctl.service.GetAll()
	if err != nil {
		return err
	}
	return c.JSON(dto.ToUserList(users))
}

func (ctl *UserController) DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user, err := ctl.service.GetUserByIdCheckAndExists(uint(id))
	if err != nil {
		return err
	}

	userCtx := context.GetUserContext(c)

	if userCtx.UserID == user.ID && userCtx.Role == "admin" {
		return &util.BadRequestError{Message: "You cannot perform this action on your own admin account."}
	}

	if err := ctl.service.Delete(user.ID); err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}
