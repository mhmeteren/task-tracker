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

// CreateUser godoc
// @Summary Register a new user
// @Description Creates a new user account. No authentication required.
// @Tags Users
// @Accept json
// @Produce json
// @Param request body dto.CreateUserRequest true "User registration data"
// @Success 201 {object} dto.ResultResponse "User successfully created"
// @Failure 400 {object} fiber.Error "Bad request (e.g., email already exists)"
// @Failure 400 {object} util.BadRequestError "Invalid request"
// @Failure 422 {object} util.ValidationError
// @Router /api/users [post]
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

// GetProfile godoc
// @Summary Get current user profile
// @Description Retrieves the profile information of the authenticated user.
// @Tags Users
// @Produce json
// @Success 200 {object} dto.UserDetail "User profile details"
// @Failure 401 {object} util.AuthError "Unauthorized, invalid or missing token"
// @Failure 404 {object} fiber.Error "User not found"
// @Security BearerAuth
// @Router /api/users/profile [get]
func (ctl *UserController) GetProfile(c *fiber.Ctx) error {
	userCtx := context.GetUserContext(c)

	user, err := ctl.service.GetProfile(userCtx.UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(dto.ToUserDetail(*user))
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Retrieves a list of all users. Requires admin role and authentication token.
// @Tags Users
// @Produce json
// @Success 200 {array} dto.UserListItem "List of users"
// @Failure 401 {object} util.AuthError "Unauthorized, invalid or missing token or permission denied"
// @Security BearerAuth
// @Router /api/users [get]
func (ctl *UserController) GetAllUsers(c *fiber.Ctx) error {
	users, err := ctl.service.GetAll()
	if err != nil {
		return err
	}
	return c.JSON(dto.ToUserList(users))
}

// DeleteUser godoc
// @Summary Delete a user by ID
// @Description Deletes a user by ID. Only users with admin role can perform this action. Admins cannot delete their own account.
// @Tags Users
// @Produce json
// @Param id path int true "User ID"
// @Success 204 "User successfully deleted"
// @Failure 400 {object} util.BadRequestError "Attempt to delete own admin account"
// @Failure 401 {object} util.AuthError "Unauthorized, invalid or missing token or permission denied"
// @Failure 404 {object} util.NotFoundError "User not found"
// @Security BearerAuth
// @Router /api/users/{id} [delete]
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
