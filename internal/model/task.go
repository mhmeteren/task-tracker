package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	UserID      uint
	TaskKey     string  `gorm:"size:10;uniqueIndex:idx_task_key_secret"`
	TaskSecret  string  `gorm:"size:10;uniqueIndex:idx_task_key_secret"`
	Name        string  `gorm:"not null;size:250"`
	Description *string `gorm:"size:500"`
	Logs        []Log   `gorm:"foreignKey:TaskID"`
}
