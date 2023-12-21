package models

import (
	"gorm.io/gorm"
)

type UsersEvents struct {
	gorm.Model
	UserID  uint `gorm:"primaryKey;autoIncrement:false"`
	EventID uint `gorm:"primaryKey;autoIncrement:false"`
}
