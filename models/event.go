package models

import (
	"github.com/pavva91/gin-gorm-rest/db"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model  `swaggerignore:"true"`
	Id          int    `json:"ID" gorm:"primary_key"`
	Category    string `json:"category" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Date        string `json:"date"`
	Time        string `json:"time"`
	Organizer   string `json:"organizer"`
	UserID      int    `json:"creator" binding:"required" swaggerignore:"true"`
}

func (h Event) CreateEvent(event *Event) (*Event, error) {
	db.GetDB().Create(&event)
	return event, nil
}

func (h Event) GetByID(id string) (*Event, error) {
	var event *Event
	db.GetDB().Where("id = ?", id).First(&event)
	return event, nil
}

func (h Event) DeleteById(id string) (*Event, error) {
	var event *Event
	db.GetDB().Where("id = ?", id).Delete(&event)
	return event, nil
}

func (h Event) SaveEvent(event *Event) (*Event, error) {
	db.GetDB().Save(&event)
	return event, nil
}
