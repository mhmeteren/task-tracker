package repository

import (
	"task-tracker/internal/model"
	"task-tracker/internal/parameter"

	"gorm.io/gorm"
)

type LogRepository interface {
	GetAllByTask(taskID uint, params *parameter.LogListParams) ([]model.Log, int64, error)

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

func (r *logRepository) GetAllByTask(taskID uint, params *parameter.LogListParams) ([]model.Log, int64, error) {

	var logs []model.Log
	var total int64

	r.db.Model(&model.Log{}).
		Where("task_id = ?", taskID).
		Count(&total)

	offset := (params.Page - 1) * params.Limit

	err := r.db.
		Where("task_id = ?", taskID).
		Order("created_at DESC").
		Limit(params.Limit).
		Offset(offset).
		Find(&logs).Error

	return logs, total, err
}
