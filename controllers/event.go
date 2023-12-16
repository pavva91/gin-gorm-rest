package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pavva91/gin-gorm-rest/auth"
	"github.com/pavva91/gin-gorm-rest/errorhandling"
	"github.com/pavva91/gin-gorm-rest/models"
	"github.com/pavva91/gin-gorm-rest/services"
	"github.com/pavva91/gin-gorm-rest/validation"
	"github.com/rs/zerolog/log"
)

var (
	Event = &event{}
)

type event struct{}

var validationUtility = new(validation.ValidationUtility)

// List godoc
//
//	@Summary		List Events
//	@Description	List all the events
//	@Tags			events
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	models.Event
//	@Router			/events [get]
//	@Schemes
func (controller event) List(c *gin.Context) {
	events, err := services.Event.ListAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error to list events", "error": err})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &events)
	c.Abort()
	return
}

// Get godoc
//
//	@Summary		Get Event
//	@Description	Get event by id
//	@Tags			events
//	@Accept			json
//	@Produce		json
//	@Param			event_id	path		int	true	"Event ID"
//	@Success		200			{object}	models.Event
//	@Failure		404			{object}	errorhandling.ErrorMessage
//	@Router			/events/{event_id} [get]
func (controller event) Get(c *gin.Context) {
	// var event models.Event
	eventId := c.Param("id")
	// if eventId != "" {
	if !validationUtility.IsEmpty(eventId) {
		// _, err := strconv.ParseUint(eventId, 10, 64)
		if !validationUtility.IsInt64(eventId) {
			r := errorhandling.SimpleErrorMessage{Message: "not valid parameter, insert valid id"}
			c.JSON(http.StatusBadRequest, r)
			c.Abort()
			return
		}
		event, err := services.Event.GetById(eventId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error to get event", "error": err})
			c.Abort()
			return
		}

		if validationUtility.IsZero(int(event.ID)) {
			r := errorhandling.SimpleErrorMessage{Message: "No event found!"}
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

// ErrorMessage represents request response with a ErrorMessage
// type ErrorMessage struct {
// 	Message string `json:"error"`
// }

// Create godoc
//
//	@Summary		Create Event
//	@Description	Create a new Event
//	@Tags			events
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.Event	true	"The new Event Values in JSON"
//	@Success		200		{object}	models.Event
//
//	@Router			/events [post]
func (controller event) Create(c *gin.Context) {
	var event models.Event

	tokenString := c.GetHeader("Authorization")

	claims, err := auth.DecodeJWT(tokenString)
	if err != nil {
		log.Err(err).Msg("unable to decode jwt token")
	}

	log.Info().Msg("username: " + claims.Username)
	user, err := services.User.GetByUsername(claims.Username)
	event.UserID = int(user.ID)

	err = c.ShouldBind(&event)
	if err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.JSON(http.StatusBadRequest, gin.H{"errors": errorhandling.NewJSONFormatter().Descriptive(verr)})
			return
		}

		// We now know that this error is not a validation error
		// probably a malformed JSON
		log.Info().Err(err).Msg("unable to bind")
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return

	}

	services.Event.Create(&event)
	c.JSON(http.StatusOK, &event)
}

// DeleteEvent godoc
//
//	@Summary		Delete Event
//	@Description	Delete event by id
//	@Tags			events
//	@Accept			json
//	@Produce		json
//	@Param			event_id	path		int	true	"Event ID"
//	@Success		200			{object}	models.Event
//
//	@Failure		404			{object}	errorhandling.ErrorMessage
//
//	@Router			/events/{event_id} [delete]
func (controller event) DeleteEvent(c *gin.Context) {
	var event models.Event
	services.Event.DeleteById(c.Param("id"))
	c.JSON(http.StatusOK, &event)
}

// SubstituteEvent godoc
//
//	@Summary		SubstituteEvent
//	@Description	Substitute the Event completely with the new JSON body
//	@Tags			events
//	@Accept			json
//	@Produce		json
//	@Param			event_id	path		int				true	"Event ID"
//	@Param			request		body		models.Event	true	"The new Event Values in JSON"
//	@Success		200			{object}	models.Event
//
//	@Failure		404			{object}	errorhandling.ErrorMessage
//
//	@Router			/events/{event_id} [put]
func (controller event) SubstituteEvent(c *gin.Context) {
	var newEvent models.Event
	eventId := c.Param("id")
	if eventId != "" {
		if !validationUtility.IsInt64(eventId) {
			errorMessage := errorhandling.SimpleErrorMessage{Message: fmt.Sprintf("not valid event id: %s - insert valid id", eventId)}
			c.JSON(http.StatusBadRequest, errorMessage)
			return
		}
		oldEvent, _ := services.Event.GetById(eventId)
		log.Info().Msg("retrieved Id: " + strconv.FormatInt(int64(oldEvent.ID), 10))
		if oldEvent.ID == 0 {
			errorMessage := errorhandling.SimpleErrorMessage{Message: "no event found!"}
			c.JSON(http.StatusNotFound, errorMessage)
			return
		}
		err := c.ShouldBind(&newEvent)
		if err != nil {
			var verr validator.ValidationErrors
			if errors.As(err, &verr) {
				errorMessage := errorhandling.ValidationErrorsMessage{Message: errorhandling.NewJSONFormatter().Descriptive(verr)}
				c.JSON(http.StatusBadRequest, errorMessage)
				return
			}

			// We now know that this error is not a validation error
			// probably a malformed JSON
			log.Info().Err(err).Msg("unable to bind")
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return

		}

		newEvent.ID = oldEvent.ID
		log.Info().Msg("retrieved Id: " + strconv.FormatInt(int64(oldEvent.ID), 10))
		services.Event.SaveEvent(&newEvent)
		c.JSON(http.StatusOK, &newEvent)
	}
}
