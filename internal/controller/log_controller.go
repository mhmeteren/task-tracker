package controller

import (
	"strconv"
	"task-tracker/internal/context"
	"task-tracker/internal/dto"
	"task-tracker/internal/model"
	"task-tracker/internal/notifier"
	"task-tracker/internal/parameter"
	"task-tracker/internal/service"
	"task-tracker/internal/util"
	"time"

	"github.com/gofiber/fiber/v2"
)

type LogController struct {
	service     service.LogService
	taskService service.TaskService
}

func NewLogController(
	service service.LogService,
	taskService service.TaskService) *LogController {
	return &LogController{service, taskService}
}

func (ctl *LogController) GetAllByTask(c *fiber.Ctx) error {
	taskID, _ := strconv.Atoi(c.Params("taskID"))

	params, err := util.BindAndSetDefaultParameters[parameter.LogListParams](c)
	if err != nil {
		return err
	}

	userCtx := context.GetUserContext(c)

	task, taskErr := ctl.taskService.GetTaskByIdAndUserCheckAndExists(uint(taskID), userCtx.UserID)
	if taskErr != nil {
		return &util.NotFoundError{Message: "Task not found"}
	}

	logs, total, err := ctl.service.GetAllByTask(task.ID, params)
	if err != nil {
		return err
	}

	return c.JSON(dto.ToPaginatedList(dto.ToLogList(logs), params.Page, params.Limit, total))
}

func (ctl *LogController) AddLog(c *fiber.Ctx) error {
	taskKey := c.Params("taskKey")
	taskSecret := c.Params("taskSecret")

	if len(taskSecret) != 10 || len(taskKey) != 10 {
		return &util.BadRequestError{Message: "Invalid keys"}
	}

	task, taskErr := ctl.taskService.FindBySecretKeyWithNotificationInfo(taskKey, taskSecret)
	if taskErr != nil {
		return &util.BadRequestError{Message: "Invalid keys"}
	}

	log := model.Log{
		TaskID:    task.ID,
		IPAddress: c.IP(),
		CreatedAt: time.Now(),
	}

	ctl.service.Create(&log)

	if task.Notification != nil { // send notification
		notifier.Enqueue(notifier.Notification{
			Message:   "Task: " + task.Name + "\nIP Address: " + log.IPAddress + "\nCreatedAt: " + log.CreatedAt.Format("2006-01-02 15:04:05.999"),
			Recipient: task.Notification.Recipient,
			Token:     task.Notification.BotToken,
			Service:   task.Notification.Service,
		})
	}

	return c.SendStatus(fiber.StatusOK)
}
