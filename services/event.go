package services

import (
	"github.com/pavva91/gin-gorm-rest/db"
	"github.com/pavva91/gin-gorm-rest/models"
)

const (
	pang = "pang"
)

var (
	EventService eventService = eventServiceImpl{}
)

type eventService interface {
	ListAllEvents() ([]models.Event, error)
	CreateEvent(event *models.Event) (*models.Event, error)
	GetById(id string) (*models.Event, error)
	DeleteById(id string) (*models.Event, error)
	SaveEvent(event *models.Event) (*models.Event, error)
}

type eventServiceImpl struct{}

func (service eventServiceImpl) ListAllEvents() ([]models.Event, error) {
	var events []models.Event
	db.GetDB().Find(&events)
	return events, nil
}

func (service eventServiceImpl) CreateEvent(event *models.Event) (*models.Event, error) {
	db.GetDB().Create(&event)
	return event, nil
}

func (service eventServiceImpl) GetById(id string) (*models.Event, error) {
	var event *models.Event
	db.GetDB().Where("id = ?", id).First(&event)
	return event, nil
}

func (service eventServiceImpl) DeleteById(id string) (*models.Event, error) {
	var event *models.Event
	db.GetDB().Where("id = ?", id).Delete(&event)
	return event, nil
}

func (service eventServiceImpl) SaveEvent(event *models.Event) (*models.Event, error) {
	db.GetDB().Save(&event)
	return event, nil
}
