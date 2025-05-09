package dto

import "task-tracker/internal/model"

type CreateTaskRequest struct {
	Name        string `json:"name" validate:"required,min=2,max=250"`
	Description string `json:"description" validate:"max=500"`
}

func (r *CreateTaskRequest) ToModel() model.Task {
	return model.Task{
		Name:        r.Name,
		Description: &r.Description,
	}
}

type UpdateTaskRequest struct {
	Name        string `json:"name" validate:"required,min=2,max=250"`
	Description string `json:"description" validate:"max=500"`
}

func (r *UpdateTaskRequest) ApplyTo(task *model.Task) {
	task.Name = r.Name
	task.Description = &r.Description
}

type TaskListItem struct {
	ID               uint                      `json:"id"`
	TaskKey          string                    `json:"task_key"`
	TaskSecret       string                    `json:"task_secret"`
	Name             string                    `json:"name"`
	Description      string                    `json:"description"`
	CreatedAt        string                    `json:"created_at"`
	TaskNotification *TaskNotificationListItem `json:"integrated_service"`
}

func ToTaskListItem(t model.Task) TaskListItem {
	return TaskListItem{
		ID:               t.ID,
		TaskKey:          t.TaskKey,
		TaskSecret:       t.TaskSecret,
		Name:             t.Name,
		Description:      *t.Description,
		CreatedAt:        t.CreatedAt.Format("2006-01-02 15:04:05"),
		TaskNotification: ToTaskNotificationListItem(t.Notification),
	}
}

func ToTaskList(t []model.Task) []TaskListItem {
	var list []TaskListItem
	for _, task := range t {
		list = append(list, ToTaskListItem(task))
	}
	return list
}
