package models

import (
	"time"

	"gorm.io/gorm"
)

type Events struct {
	gorm.Model
	Title            string
	Description      string
	Date             time.Time
	Time             string
	LocationID       uint
	CategoryID       uint
	UsersEvents      []UsersEvents `gorm:"foreignKey:EventID"`
	OrganizersEvents []UsersEvents `gorm:"foreignKey:EventID"`
}
