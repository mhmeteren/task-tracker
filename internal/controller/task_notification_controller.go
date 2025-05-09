package controller

import (
	"strconv"
	"task-tracker/internal/context"
	"task-tracker/internal/dto"
	"task-tracker/internal/notifier"
	"task-tracker/internal/service"
	"task-tracker/internal/util"

	"github.com/gofiber/fiber/v2"
)

type TaskNotificationController struct {
	service     service.TaskNotificationService
	taskService service.TaskService
}

func NewTaskNotificationController(
	service service.TaskNotificationService,
	taskService service.TaskService) *TaskNotificationController {
	return &TaskNotificationController{service, taskService}
}

func (ctl *TaskNotificationController) NotificationServiceIntegration(c *fiber.Ctx) error {
	req, err := util.BindAndValidate[dto.CreateTaskNotificationRequest](c)
	if err != nil {
		return err
	}

	userCtx := context.GetUserContext(c)
	task, err := ctl.taskService.GetTaskByIdAndUserCheckAndExists(req.TaskID, userCtx.UserID)
	if err != nil {
		return err
	}

	_, err = ctl.service.FindByTask(task.ID)
	if err == nil {
		return &util.NotFoundError{Message: "An integration already exists for this task."}
	}

	taskNoti := req.ToModel()
	taskNoti.TaskID = task.ID
	if err := ctl.service.CreateNewTask(&taskNoti); err != nil {
		return err
	}

	//Send Test Notification
	notifier.Enqueue(notifier.Notification{
		Message: "Test message: your configuration is valid.",
		ChatID:  taskNoti.ChatID,
		Token:   taskNoti.BotToken,
		Service: taskNoti.Service,
	})

	return c.Status(fiber.StatusCreated).JSON(dto.ResultResponse{Message: req.Service + " integration is successfully. Check your log chat"})
}

func (ctl *TaskNotificationController) DeleteNotificationServiceInformation(c *fiber.Ctx) error {
	task_id, _ := strconv.Atoi(c.Params("taskID"))

	userCtx := context.GetUserContext(c)

	task, err := ctl.taskService.GetTaskByIdAndUserCheckAndExists(uint(task_id), userCtx.UserID)
	if err != nil {
		return err
	}

	taskNotification, err := ctl.service.FindByTask(task.ID)
	if err != nil {
		return &util.NotFoundError{Message: "integration information not found"}
	}

	if err := ctl.service.Delete(taskNotification.ID); err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}
