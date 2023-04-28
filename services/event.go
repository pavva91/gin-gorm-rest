package services

import (
	"github.com/pavva91/gin-gorm-rest/db"
	"github.com/pavva91/gin-gorm-rest/models"
	"github.com/rs/zerolog/log"
)

const (
	pang = "pang"
)

var (
	EventService eventService = eventServiceImpl{}
)

type eventService interface {
	ListAllEvents() ([]models.Event, error)
}

type eventServiceImpl struct{}

func (service eventServiceImpl) ListAllEvents() ([]models.Event, error) {
	log.Info().Msg("I'm inside EventService")
	var events []models.Event
	db.GetDB().Find(&events)
	return events, nil
}
