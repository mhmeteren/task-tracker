package controller

import (
	"strconv"
	"task-tracker/internal/context"
	"task-tracker/internal/dto"
	"task-tracker/internal/service"
	"task-tracker/internal/util"

	"github.com/gofiber/fiber/v2"
)

type TaskController struct {
	service service.TaskService
}

func NewTaskController(service service.TaskService) *TaskController {
	return &TaskController{service}
}

// GetAllByLoggedUser godoc
// @Summary Get all tasks for the authenticated user
// @Description Returns a list of tasks belonging to the logged-in user
// @Tags Tasks
// @Produce json
// @Success 200 {array} dto.TaskListItem
// @Failure 401 {object} util.AuthError "Unauthorized, invalid or missing token"
// @Security BearerAuth
// @Router /api/tasks [get]
func (ctl *TaskController) GetAllByLoggedUser(c *fiber.Ctx) error {
	userCtx := context.GetUserContext(c)

	tasks, err := ctl.service.GetAllByUser(userCtx.UserID)
	if err != nil {
		return err
	}
	return c.JSON(dto.ToTaskList(tasks))
}

// CreateNewTask godoc
// @Summary Create a new task
// @Description Creates a new task for the authenticated user
// @Tags Tasks
// @Accept json
// @Produce json
// @Param request body dto.CreateTaskRequest true "Task creation payload"
// @Success 201 {object} dto.ResultResponse "Task successfully created"
// @Failure 400 {object} util.BadRequestError "Invalid request"
// @Failure 401 {object} util.AuthError "Unauthorized, invalid or missing token"
// @Failure 422 {object} util.ValidationError
// @Security BearerAuth
// @Router /api/tasks [post]
func (ctl *TaskController) CreateNewTask(c *fiber.Ctx) error {
	req, err := util.BindAndValidate[dto.CreateTaskRequest](c)
	if err != nil {
		return err
	}

	task := req.ToModel()

	userCtx := context.GetUserContext(c)
	task.UserID = userCtx.UserID

	if err := ctl.service.CreateNewTask(&task); err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ResultResponse{Message: "Task successfully created"})
}

// UpdateTask godoc
// @Summary Update an existing task
// @Description Updates a task by ID for the authenticated user
// @Tags Tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param request body dto.UpdateTaskRequest true "Task update payload"
// @Success 204 {object} dto.ResultResponse "Task successfully updated"
// @Failure 400 {object} util.BadRequestError "Invalid request"
// @Failure 401 {object} util.AuthError "Unauthorized, invalid or missing token"
// @Failure 404 {object} util.NotFoundError "Task not found"
// @Failure 422 {object} util.ValidationError
// @Security BearerAuth
// @Router /api/tasks/{id} [put]
func (ctl *TaskController) UpdateTask(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	req, err := util.BindAndValidate[dto.UpdateTaskRequest](c)
	if err != nil {
		return err
	}

	userCtx := context.GetUserContext(c)

	task, err := ctl.service.GetTaskByIdAndUserCheckAndExists(uint(id), userCtx.UserID)
	if err != nil {
		return err
	}

	req.ApplyTo(task)

	if err := ctl.service.Update(task); err != nil {
		return err
	}

	return c.Status(fiber.StatusNoContent).JSON(dto.ResultResponse{Message: "Task successfully updated"})
}

// DeleteTask godoc
// @Summary Delete a task
// @Description Deletes a task by ID for the authenticated user
// @Tags Tasks
// @Produce json
// @Param id path int true "Task ID"
// @Success 204 "Task successfully deleted"
// @Failure 401 {object} util.AuthError "Unauthorized, invalid or missing token"
// @Failure 404 {object} util.NotFoundError "Task not found"
// @Security BearerAuth
// @Router /api/tasks/{id} [delete]
func (ctl *TaskController) DeleteTask(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	userCtx := context.GetUserContext(c)

	task, err := ctl.service.GetTaskByIdAndUserCheckAndExists(uint(id), userCtx.UserID)
	if err != nil {
		return err
	}

	if err := ctl.service.Delete(task.ID); err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}
