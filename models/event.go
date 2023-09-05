package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model  `swaggerignore:"true"`
	Category    string    `json:"category" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Date        time.Time `json:"date"`
	Time        string    `json:"time"`
	Organizer   string    `json:"organizer"`
	UserID      int       `json:"creator" binding:"required" swaggerignore:"true"`
}
