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

func (ctl *TaskController) GetAllByLoggedUser(c *fiber.Ctx) error {
	userCtx := context.GetUserContext(c)

	tasks, err := ctl.service.GetAllByUser(userCtx.UserID)
	if err != nil {
		return err
	}
	return c.JSON(dto.ToTaskList(tasks))
}

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
