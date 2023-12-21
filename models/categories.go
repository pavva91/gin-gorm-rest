package models

import (
	"gorm.io/gorm"
)

type Categories struct {
	gorm.Model
	Events  []Events `gorm:"foreignKey:CategoryID"`
}
