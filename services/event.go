package services

import (
	"github.com/pavva91/gin-gorm-rest/db"
	"github.com/pavva91/gin-gorm-rest/models"
)

var (
	Event eventService = event{}
)

type eventService interface {
	ListAllEvents() ([]models.Event, error)
	Create(e *models.Event) (*models.Event, error)
	GetById(id string) (*models.Event, error)
	DeleteById(id string) (*models.Event, error)
	SaveEvent(e *models.Event) (*models.Event, error)
}

type event struct{}

func (s event) ListAllEvents() ([]models.Event, error) {
	var events []models.Event
	// TODO: Use Repository
	db.DbOrm.GetDB().Find(&events)
	return events, nil
}

func (s event) Create(e *models.Event) (*models.Event, error) {
	db.DbOrm.GetDB().Create(&e)
	return e, nil
}

func (s event) GetById(id string) (*models.Event, error) {
	var event *models.Event
	db.DbOrm.GetDB().Where("id = ?", id).First(&event)
	return event, nil
}

func (s event) DeleteById(id string) (*models.Event, error) {
	var event *models.Event
	db.DbOrm.GetDB().Where("id = ?", id).Delete(&event)
	return event, nil
}

func (s event) SaveEvent(e *models.Event) (*models.Event, error) {
	db.DbOrm.GetDB().Save(&e)
	return e, nil
}
