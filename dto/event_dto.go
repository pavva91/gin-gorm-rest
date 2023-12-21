package dto

import (
	"time"

	"github.com/pavva91/gin-gorm-rest/models"
)

type EventDTO struct {
	ID          uint      `json:"eventID"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Time        string    `json:"time"`
	LocationID  uint      `json:"locationID"`
}

func (dto *EventDTO) ToModel() *models.Events {
	var model *models.Events
	model.Title = dto.Title
	model.Description = dto.Description
	model.Date = dto.Date
	model.Time = dto.Time
	model.LocationID = dto.LocationID
	return model
}

func (dto *EventDTO) ToDTO(model models.Events) {
	dto.ID = model.ID
	dto.Title = model.Title
	dto.Description = model.Description
	dto.Date = model.Date
	dto.Time = model.Time
	dto.LocationID = model.LocationID
}

func (dto *EventDTO) ToDTOs(models []models.Events) (dtos []EventDTO) {
	dtos = make([]EventDTO, len(models))
	for i, v := range models {
		dtos[i].ToDTO(v)
	}
	return dtos
}
