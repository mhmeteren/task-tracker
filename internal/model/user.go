package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name                  string `gorm:"not null"`
	Email                 string `gorm:"uniqueIndex;not null"`
	Password              string `gorm:"not null"`
	UserKey               string `gorm:"uniqueIndex;size:10"`
	RefreshToken          string
	RefreshTokenExpiresAt time.Time
	RoleID                uint
	Role                  Role   `gorm:"foreignKey:RoleID"`
	Tasks                 []Task `gorm:"foreignKey:UserID"`
}
