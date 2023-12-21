package models

import (
	"gorm.io/gorm"
)

type Locations struct {
	gorm.Model
	ZipCode uint
	State   uint
	City    string
	Street  string
	Number  uint
	Events  []Events `gorm:"foreignKey:LocationID"`
}
