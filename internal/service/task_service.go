package service

import (
	"task-tracker/internal/model"
	"task-tracker/internal/repository"
	"task-tracker/internal/util"
)

type TaskService interface {
	GetAllByUser(userID uint) ([]model.Task, error)
	GetTaskByIdAndUserCheckAndExists(id uint, userID uint) (*model.Task, error)
	CreateNewTask(task *model.Task) error
	Update(task *model.Task) error
	Delete(id uint) error
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{repo}
}

func (service *taskService) GetAllByUser(userID uint) ([]model.Task, error) {
	return service.repo.GetAllByUser(userID)
}

func (service *taskService) CreateNewTask(task *model.Task) error {
	task.TaskKey = util.GenerateKey(10)
	return service.repo.Create(task)
}

func (s *taskService) Update(task *model.Task) error {
	return s.repo.Update(task)
}

func (s *taskService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *taskService) GetTaskByIdAndUserCheckAndExists(id uint, userID uint) (*model.Task, error) {
	task, err := s.repo.FindByID(id)
	if err != nil || userID != task.UserID {
		return nil, &util.NotFoundError{Message: "Task not found"}
	}

	return task, nil
}
