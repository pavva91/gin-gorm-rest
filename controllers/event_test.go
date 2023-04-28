package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pavva91/gin-gorm-rest/models"
	"github.com/pavva91/gin-gorm-rest/services"
	"github.com/stretchr/testify/assert"
)

type eventServiceMock struct {
	listEventsFn func() ([]models.Event, error)
}

func (mock eventServiceMock) ListAllEvents() ([]models.Event, error) {
	return mock.listEventsFn()
}

func Test_ListEvents_Error_Error(t *testing.T) {
	expectedHttpStatus := http.StatusInternalServerError
	expectedHttpBody := "{\"error\":{},\"message\":\"Error to list events\"}"

	serviceMock := eventServiceMock{}
	serviceMock.listEventsFn = func() ([]models.Event, error) {
		return nil, errors.New("error executing ping")
	}
	services.EventService = serviceMock

	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)

	EventController.ListE(context)

	actualHttpStatus := context.Writer.Status()
	actualHttpBody := response.Body.String()

	assert.Equal(t, actualHttpStatus, expectedHttpStatus)
	assert.Equal(t, actualHttpBody, expectedHttpBody)
}

func Test_ListEvents_Empty_Empty(t *testing.T) {
	emptyEventList := []models.Event{}

	expectedHttpStatus := http.StatusOK
	expectedHttpBody := "[]"

	serviceMock := eventServiceMock{}
	serviceMock.listEventsFn = func() ([]models.Event, error) {
		return emptyEventList, nil
	}
	services.EventService = serviceMock

	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)

	EventController.ListE(context)

	actualHttpStatus := context.Writer.Status()
	actualHttpBody := response.Body.String()

	assert.Equal(t, actualHttpStatus, expectedHttpStatus)
	assert.Equal(t, actualHttpBody, expectedHttpBody)
}
