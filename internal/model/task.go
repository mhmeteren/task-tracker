package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	UserID      uint
	TaskKey     string `gorm:"uniqueIndex;size:32"`
	Name        string `gorm:"size:250"`
	Description string `gorm:"size:500"`
	Logs        []Log
}
