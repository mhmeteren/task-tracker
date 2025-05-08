package repository

import (
	"task-tracker/internal/model"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetAllByUser(userID uint) ([]model.Task, error)
	FindByID(id uint) (*model.Task, error)
	FindBySecretKey(taskKey, taskSecret string) (*model.Task, error)

	Create(task *model.Task) error
	Update(task *model.Task) error
	Delete(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) GetAllByUser(userID uint) ([]model.Task, error) {
	var tasks []model.Task
	err := r.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&tasks).Error

	return tasks, err
}

func (r *taskRepository) FindByID(id uint) (*model.Task, error) {
	var task model.Task
	if err := r.db.First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *taskRepository) FindBySecretKey(taskKey, taskSecret string) (*model.Task, error) {
	var task model.Task
	err := r.db.
		Where("task_key = ? AND task_secret = ?", taskKey, taskSecret).
		First(&task).Error
	return &task, err
}

func (r *taskRepository) Create(task *model.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) Update(task *model.Task) error {
	return r.db.Save(task).Error
}

func (r *taskRepository) Delete(id uint) error {
	return r.db.Delete(&model.Task{}, id).Error
}
