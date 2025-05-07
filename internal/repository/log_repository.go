package repository

import (
	"task-tracker/internal/model"

	"gorm.io/gorm"
)

type LogRepository interface {
	GetAllByTask(taskID uint) ([]model.Log, error)

	Create(user *model.Log) error
}

type logRepository struct {
	db *gorm.DB
}

func NewLogRepository(db *gorm.DB) LogRepository {
	return &logRepository{db}
}

func (r *logRepository) Create(user *model.Log) error {
	return r.db.Create(user).Error
}

func (r *logRepository) GetAllByTask(taskID uint) ([]model.Log, error) {
	var logs []model.Log
	err := r.db.Where("task_id = ?", taskID).Find(&logs).Error

	return logs, err
}
