package service

import (
	"task-tracker/internal/model"
	"task-tracker/internal/repository"
)

type TaskNotificationService interface {
	FindByTask(task_id uint) (*model.TaskNotification, error)

	CreateNewTask(task *model.TaskNotification) error
	Update(task *model.TaskNotification) error
	Delete(id uint) error
}

type taskNotificationService struct {
	repo repository.TaskNotificationRepository
}

func NewTaskNotificationService(repo repository.TaskNotificationRepository) TaskNotificationService {
	return &taskNotificationService{repo}
}

func (s *taskNotificationService) FindByTask(task_id uint) (*model.TaskNotification, error) {
	return s.repo.FindByTask(task_id)
}

func (service *taskNotificationService) CreateNewTask(task *model.TaskNotification) error {
	return service.repo.Create(task)
}

func (s *taskNotificationService) Update(task *model.TaskNotification) error {
	return s.repo.Update(task)
}

func (s *taskNotificationService) Delete(id uint) error {
	return s.repo.Delete(id)
}
