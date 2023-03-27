package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pavva91/gin-gorm-rest/formatter"
	"github.com/pavva91/gin-gorm-rest/models"
	"github.com/rs/zerolog/log"
)

type EventController struct{}

var eventModel = new(models.Event)

// ListEvents godoc
//
//	@Summary		List Events
//	@Description	List all the events
//	@Tags			events
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	models.Event
//	@Router			/events [get]
//	@Schemes
func (ec EventController) ListEvents(c *gin.Context) {
	events, err := eventModel.ListAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to list events", "error": err})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &events)
	c.Abort()
	return
}

// ListEvents godoc
//
//	@Summary		Get Event
//	@Description	Get event by id
//	@Tags			events
//	@Accept			json
//	@Produce		json
//	@Param event_id   path int true "Event ID"
//	@Success		200	{object}	models.Event
//	@Failure		404	{object}	message
//	@Router			/events/{event_id} [get]
func GetEvent(c *gin.Context) {
	// var event models.Event
	eventId := c.Param("id")
	if eventId != "" {
		_, err := strconv.ParseUint(eventId, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Not valid parameter, Insert valid id"})
			return
		}
		event, err := eventModel.GetByID(eventId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user", "error": err})
			c.Abort()
			return
		}

		if event.Id == 0 {
			r := message{"No event found!"}
			// c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No event found!"})
			c.JSON(http.StatusNotFound, r)
			return
		} else {
			// c.JSON(http.StatusOK, gin.H{"message": "Event founded!", "event": event})
			c.JSON(http.StatusOK, event)
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
	return
}

// message represents request response with a message
type message struct {
	Message string `json:"message"`
}

// CreateEvent godoc
//
//		@Summary		Create Event
//		@Description	Create a new Event
//		@Tags			events
//		@Accept			json
//		@Produce		json
//	 @Param			request body models.Event true "The new Event Values in JSON"
//		@Success		200	{object}	models.Event
//
//		@Router			/events [post]
func CreateEvent(c *gin.Context) {
	var event models.Event
	// err := c.BindJSON(&event)
	err := c.ShouldBind(&event)
	if err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.JSON(http.StatusBadRequest, gin.H{"errors": formatter.NewJSONFormatter().Descriptive(verr)})
			return
		}

		// We now know that this error is not a validation error
		// probably a malformed JSON
		log.Info().Err(err).Msg("unable to bind")
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return

	}

	// c.BindJSON(&event)
	eventModel.CreateEvent(&event)
	c.JSON(http.StatusOK, &event)
}

// func Simple(verr validator.ValidationErrors) map[string]string {
// 	errs := make(map[string]string)
//
// 	for _, f := range verr {
// 		err := f.ActualTag()
// 		if f.Param() != "" {
// 			err = fmt.Sprintf("%s=%s", err, f.Param())
// 		}
// 		errs[f.Field()] = err
// 	}
//
// 	return errs
// }

// DeleteEvent godoc
//
//	@Summary		Delete Event
//	@Description	Delete event by id
//	@Tags			events
//	@Accept			json
//	@Produce		json
//	@Param event_id   path int true "Event ID"
//	@Success		200	{object}	models.Event
//
//	@Failure		404	{object}	message
//
//	@Router			/events/{event_id} [delete]
func DeleteEvent(c *gin.Context) {
	var event models.Event
	eventModel.DeleteById(c.Param("id"))
	c.JSON(http.StatusOK, &event)
}

// SubstituteEvent godoc
//
//		@Summary		SubstituteEvent
//		@Description	Substitute the Event completely with the new JSON body
//		@Tags			events
//		@Accept			json
//		@Produce		json
//		@Param event_id   path int true "Event ID"
//	 @Param			request body models.Event true "The new Event Values in JSON"
//		@Success		200	{object}	models.Event
//
//		@Failure		404	{object}	message
//
//		@Router			/events/{event_id} [put]
func SubstituteEvent(c *gin.Context) {
	var event models.Event
	eventId := c.Param("id")
	if eventId != "" {
		_, id_err := strconv.ParseUint(eventId, 10, 64)
		if id_err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Not valid parameter, Insert valid id"})
			return
		}
		eventModel.GetByID(c.Param("id"))
		if event.Id == 0 {
			r := message{"No event found!"}
			// c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No event found!"})
			c.JSON(http.StatusNotFound, r)
			return
		}
		err := c.ShouldBind(&event)
		if err != nil {
			var verr validator.ValidationErrors
			if errors.As(err, &verr) {
				c.JSON(http.StatusBadRequest, gin.H{"errors": formatter.NewJSONFormatter().Descriptive(verr)})
				return
			}

			// We now know that this error is not a validation error
			// probably a malformed JSON
			log.Info().Err(err).Msg("unable to bind")
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return

		}
		eventModel.SaveEvent(&event)
		c.JSON(http.StatusOK, &event)
	}
}
