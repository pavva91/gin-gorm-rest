package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pavva91/gin-gorm-rest/models"
	"github.com/pavva91/gin-gorm-rest/services"
	"github.com/pavva91/gin-gorm-rest/stubs"
	"github.com/stretchr/testify/assert"
)

func Test_GetByID_EmptyId_400BadRequest(t *testing.T) {
	// Mocks
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)

	// Expected Result
	expectedHttpStatus := http.StatusBadRequest
	expectedHttpBody := "{\"error\":\"empty id\"}"

	// Call functio to test
	UserController.GetUser(context)

	// Check Values
	actualHttpStatus := context.Writer.Status()
	actualHttpBody := response.Body.String()

	assert.Equal(t, expectedHttpStatus, actualHttpStatus)
	assert.Equal(t, expectedHttpBody, actualHttpBody)
}

func Test_GetByID_NoIntId_400BadRequest(t *testing.T) {
	// Mocks
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.AddParam("id", "not an int")

	// Expected Result
	expectedHttpStatus := http.StatusBadRequest
	expectedError := "Not valid parameter, Insert valid id"
	expectedHttpBody := "{\"error\":\"" + expectedError + "\"}"

	// Call function to test
	UserController.GetUser(context)

	// Check Values
	actualHttpStatus := context.Writer.Status()
	actualHttpBody := response.Body.String()

	assert.Equal(t, expectedHttpStatus, actualHttpStatus)
	assert.Equal(t, expectedHttpBody, actualHttpBody)
	assert.Contains(t, actualHttpBody, expectedError)
}
func Test_GetByID_InternalErrorGetById_500InternalServerError(t *testing.T) {
	// Mocks
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.AddParam("id", "0")

	// Stubs
	UserServiceStub := stubs.UserServiceStub{}
	UserServiceStub.GetByIDFn = func() (*models.User, error) {
		return nil, errors.New("error stub")
	}
	services.UserService = UserServiceStub

	// Expected Result
	expectedHttpStatus := http.StatusInternalServerError
	expectedError := "Error to get user"
	expectedHttpBody := "{\"error\":\"" + expectedError + "\"}"

	// Call function to test
	UserController.GetUser(context)

	// Check Values
	actualHttpStatus := context.Writer.Status()
	actualHttpBody := response.Body.String()

	assert.Equal(t, expectedHttpStatus, actualHttpStatus)
	assert.Equal(t, expectedHttpBody, actualHttpBody)
	assert.Contains(t, actualHttpBody, expectedError)
}

func Test_GetByID_NotFoundId_404NotFound(t *testing.T) {

	// Mocks
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.AddParam("id", "0")

	// Stubs
	var userStub models.User
	userStub.ID = 0
	UserServiceStub := stubs.UserServiceStub{}
	UserServiceStub.GetByIDFn = func() (*models.User, error) {
		return &userStub, nil
	}
	services.UserService = UserServiceStub

	// Expected Values
	expectedHttpStatus := http.StatusNotFound
	expectedError := "No user found"
	expectedHttpBody := "{\"error\":\"" + expectedError + "\"}"

	// Call function to test
	UserController.GetUser(context)

	// Check Values
	actualHttpStatus := context.Writer.Status()
	actualHttpBody := response.Body.String()
	assert.Contains(t, actualHttpBody, expectedError)

	assert.Equal(t, expectedHttpStatus, actualHttpStatus)
	assert.Equal(t, expectedHttpBody, actualHttpBody)
}

func Test_GetByID_FoundId_200ReturnUser(t *testing.T) {

	// Mocks
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.AddParam("id", "1")

	// Stubs
	var userStub models.User
	userStub.ID = 1
	userStub.Name = "Kurt"
	userStub.Username = "user1234"
	userStub.Password = "encrypted"

	UserServiceStub := stubs.UserServiceStub{}
	UserServiceStub.GetByIDFn = func() (*models.User, error) {
		return &userStub, nil
	}
	services.UserService = UserServiceStub

	// Expected Values
	expectedHttpStatus := http.StatusOK
	expectedHttpBody := "{\"ID\":1,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"DeletedAt\":null,\"name\":\"Kurt\",\"username\":\"user1234\",\"email\":\"\",\"password\":\"encrypted\",\"Events\":null}"

	// Call function to test
	UserController.GetUser(context)

	// Check Values
	actualHttpStatus := context.Writer.Status()
	actualHttpBody := response.Body.String()

	assert.Equal(t, expectedHttpStatus, actualHttpStatus)
	assert.Equal(t, expectedHttpBody, actualHttpBody)
}
