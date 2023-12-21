package models

import (
	"gorm.io/gorm"
)

type OrganizersEvents struct {
	gorm.Model
	OrganizerID uint `gorm:"primaryKey;autoIncrement:false"`
	EventID     uint `gorm:"primaryKey;autoIncrement:false"`
}
