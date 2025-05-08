package service

import (
	"task-tracker/internal/model"
	"task-tracker/internal/parameter"
	"task-tracker/internal/repository"
)

type LogService interface {
	GetAllByTask(taskID uint, params *parameter.LogListParams) ([]model.Log, int64, error)

	Create(log *model.Log) error
}

type logService struct {
	repo repository.LogRepository
}

func NewLogService(repo repository.LogRepository) LogService {
	return &logService{repo}
}

func (s *logService) GetAllByTask(taskID uint, params *parameter.LogListParams) ([]model.Log, int64, error) {
	return s.repo.GetAllByTask(taskID, params)
}

func (s *logService) Create(log *model.Log) error {
	return s.repo.Create(log)
}
