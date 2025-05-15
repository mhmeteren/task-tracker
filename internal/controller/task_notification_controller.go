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

// NotificationServiceIntegration godoc
// @Summary Integrate a notification service to a task
// @Description Adds a notification integration (Telegram, Slack, Discord, etc.) to a user's task and sends a test notification
// @Tags Task Notifications
// @Accept json
// @Produce json
// @Param request body dto.CreateTaskNotificationRequest true "Notification integration details"
// @Success 201 {object} dto.ResultResponse "Integration created successfully and test notification sent"
// @Failure 400 {object} util.BadRequestError "Invalid request"
// @Failure 401 {object} util.AuthError "Unauthorized, invalid or missing token"
// @Failure 404 {object} util.NotFoundError "Task not found or integration already exists"
// @Failure 422 {object} util.ValidationError
// @Security BearerAuth
// @Router /api/task-notifications [post]
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
		Message:   "Test message: your configuration is valid.",
		Recipient: taskNoti.Recipient,
		Token:     taskNoti.BotToken,
		Service:   taskNoti.Service,
	})

	return c.Status(fiber.StatusCreated).JSON(dto.ResultResponse{Message: req.Service + " integration is successfully. Check your log chat"})
}

// DeleteNotificationServiceInformation godoc
// @Summary Delete notification integration for a task
// @Description Deletes the notification service integration information for the specified task of the authenticated user
// @Tags Task Notifications
// @Produce json
// @Param taskID path int true "Task ID"
// @Success 204 "Notification integration successfully deleted"
// @Failure 401 {object} util.AuthError "Unauthorized, invalid or missing token"
// @Failure 404 {object} util.NotFoundError "Integration information not found"
// @Security BearerAuth
// @Router /api/task-notifications/{taskID} [delete]
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
