package dto

import "task-tracker/internal/model"

type CreateTaskNotificationRequest struct {
	TaskID    uint   `json:"task_id" validate:"required"`
	Service   string `json:"service" validate:"required,oneof=telegram slack"`
	BotToken  string `json:"bot_token" validate:"required"`
	Recipient string `json:"recipient" validate:"required"`
}

func (r *CreateTaskNotificationRequest) ToModel() model.TaskNotification {
	return model.TaskNotification{
		Service:   r.Service,
		BotToken:  r.BotToken,
		Recipient: r.Recipient,
	}
}

type TaskNotificationListItem struct {
	Service   string `json:"service"`
	CreatedAt string `json:"created_at"`
}

func ToTaskNotificationListItem(t *model.TaskNotification) *TaskNotificationListItem {

	if t == nil {
		return nil
	}

	return &TaskNotificationListItem{
		Service:   t.Service,
		CreatedAt: t.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
