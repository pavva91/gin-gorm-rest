package models

import (
	"github.com/pavva91/gin-gorm-rest/db"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model  `swaggerignore:"true"`
	Id          int    `json:"ID" gorm:"primary_key"`
	Category    string `json:"category"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Date        string `json:"date"`
	Time        string `json:"time"`
	Organizer   string `json:"organizer"`
}

func (h Event) ListAllEvents() ([]Event, error) {
	var events []Event
	db.GetDB().Find(&events)
	return events, nil
}

func (h Event) CreateEvent(*Event) (*Event, error) {
	var event *Event
	db.GetDB().Create(&event)
	return event, nil
}

func (h Event) GetByID(id string) (*Event, error) {
	var event *Event
	db.GetDB().Where("id = ?", id).First(&event)
	return event, nil
}
