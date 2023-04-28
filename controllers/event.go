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
	EventController = eventController{}
)

type eventController struct{}

var eventModel = new(models.Event)
var validationController = new(validation.ValidationController)

func (controller eventController) ListE(context *gin.Context) {
	events, err := services.EventService.ListAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error to list events", "error": err})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, &events)
	context.Abort()
	return
}

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

// func (ec eventController) ListEvents(c *gin.Context) {
// 	events, err := eventModel.ListAllEvents()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to list events", "error": err})
// 		c.Abort()
// 		return
// 	}
// 	c.JSON(http.StatusOK, &events)
// 	c.Abort()
// 	return
// }

// GetEvent godoc
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
func GetEvent(c *gin.Context) {
	// var event models.Event
	eventId := c.Param("id")
	// if eventId != "" {
	if !validationController.IsEmpty(eventId) {
		// _, err := strconv.ParseUint(eventId, 10, 64)
		if !validationController.IsInt64(eventId) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Not valid parameter, Insert valid id"})
			return
		}
		event, err := eventModel.GetByID(eventId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user", "error": err})
			c.Abort()
			return
		}

		if validationController.IsZero(event.Id) {
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

// CreateEvent godoc
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
func CreateEvent(c *gin.Context) {
	var event models.Event
	var userModel models.User

	tokenString := c.GetHeader("Authorization")

	claims, err := auth.DecodeJWT(tokenString)
	if err != nil {
		log.Err(err).Msg("Unable to Decode JWT Token")
	}

	log.Info().Msg("Username: " + claims.Username)
	user, err := userModel.GetByUsername(claims.Username)
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
//	@Param			event_id	path		int	true	"Event ID"
//	@Success		200			{object}	models.Event
//
//	@Failure		404			{object}	errorhandling.ErrorMessage
//
//	@Router			/events/{event_id} [delete]
func DeleteEvent(c *gin.Context) {
	var event models.Event
	eventModel.DeleteById(c.Param("id"))
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
func SubstituteEvent(c *gin.Context) {
	var newEvent models.Event
	eventId := c.Param("id")
	if eventId != "" {
		if !validationController.IsInt64(eventId) {
			errorMessage := errorhandling.SimpleErrorMessage{Message: fmt.Sprintf("Not valid event id: %s - Insert valid id", eventId)}
			c.JSON(http.StatusBadRequest, errorMessage)
			return
		}
		oldEvent, _ := eventModel.GetByID(eventId)
		log.Info().Msg("retrieved Id: " + strconv.FormatInt(int64(oldEvent.Id), 10))
		if oldEvent.Id == 0 {
			errorMessage := errorhandling.SimpleErrorMessage{Message: "No event found!"}
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

		newEvent.Id = oldEvent.Id
		log.Info().Msg("retrieved Id: " + strconv.FormatInt(int64(oldEvent.Id), 10))
		eventModel.SaveEvent(&newEvent)
		c.JSON(http.StatusOK, &newEvent)
	}
}
