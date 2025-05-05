package model

import "time"

type Log struct {
	ID        uint `gorm:"primaryKey"`
	TaskID    uint
	IPAddress string `gorm:"size:45"`
	CreatedAt time.Time
}
