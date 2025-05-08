package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name                  string  `gorm:"not null"`
	Email                 string  `gorm:"uniqueIndex;not null"`
	Password              string  `gorm:"not null"`
	RefreshToken          *string `gorm:"uniqueIndex"`
	RefreshTokenExpiresAt time.Time
	RoleID                uint
	Role                  Role   `gorm:"foreignKey:RoleID"`
	Tasks                 []Task `gorm:"foreignKey:UserID"`
}
