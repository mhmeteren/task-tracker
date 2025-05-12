package model

import "time"

type TaskNotification struct {
	ID        uint   `gorm:"primaryKey"`
	Service   string `gorm:"not null"`
	BotToken  string `gorm:"not null"`
	Recipient string `gorm:"not null"`
	CreatedAt time.Time

	TaskID uint `gorm:"not null;unique"`
}
