package repository

import (
	"task-tracker/internal/model"

	"gorm.io/gorm"
)

type TaskNotificationRepository interface {
	FindByTask(task_id uint) (*model.TaskNotification, error)

	Create(task *model.TaskNotification) error
	Update(task *model.TaskNotification) error
	Delete(id uint) error
}

type taskNotificationRepository struct {
	db *gorm.DB
}

func NewTaskNotificationRepository(db *gorm.DB) TaskNotificationRepository {
	return &taskNotificationRepository{db}
}

func (r *taskNotificationRepository) FindByTask(task_id uint) (*model.TaskNotification, error) {
	var taskNotification model.TaskNotification
	if err := r.db.Where("task_id = ?", task_id).First(&taskNotification).Error; err != nil {
		return nil, err
	}
	return &taskNotification, nil
}

func (r *taskNotificationRepository) Create(task *model.TaskNotification) error {
	return r.db.Create(task).Error
}

func (r *taskNotificationRepository) Update(task *model.TaskNotification) error {
	return r.db.Save(task).Error
}

func (r *taskNotificationRepository) Delete(id uint) error {
	return r.db.Delete(&model.TaskNotification{}, id).Error
}
