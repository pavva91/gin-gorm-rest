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

type userServiceMock struct {
	listUsersFn     func() ([]models.User, error)
	getByIDFn       func() (*models.User, error)
	getByUsernameFn func() (*models.User, error)
	updateFn        func() (*models.User, error)
	deleteFn        func() (*models.User, error)
}

func (mock userServiceMock) ListUsers() ([]models.User, error) {
	return mock.listUsersFn()
}

func (mock userServiceMock) GetByID(id string) (*models.User, error) {
	return mock.getByIDFn()
}

func (mock userServiceMock) GetByUsername(username string) (*models.User, error) {
	return mock.getByUsernameFn()
}

func (mock userServiceMock) Update(user *models.User) (*models.User, error) {
	return mock.updateFn()
}

func (mock userServiceMock) Delete(id string) (*models.User, error) {
	return mock.deleteFn()
}

func Test_GetByID_EmptyId_400BadRequest(t *testing.T) {
	expectedHttpStatus := http.StatusBadRequest
	expectedHttpBody := "{\"error\":\"empty id\"}"

	// userServiceMock := userServiceMock{}
	// userServiceMock.getByIDFn = func() (*models.User, error) {
	// 	return nil, errors.New("error executing ping")
	// }
	// services.UserService = userServiceMock

	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)

	UserController.GetUser(context)

	actualHttpStatus := context.Writer.Status()
	actualHttpBody := response.Body.String()

	assert.Equal(t, expectedHttpStatus, actualHttpStatus)
	assert.Equal(t, expectedHttpBody, actualHttpBody)
}

func Test_GetByID_NoIntId_400BadRequest(t *testing.T) {
	expectedHttpStatus := http.StatusBadRequest
	expectedError := "Not valid parameter, Insert valid id"
	expectedHttpBody := "{\"error\":\"" + expectedError + "\"}"

	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.AddParam("id", "not an int")

	UserController.GetUser(context)

	actualHttpStatus := context.Writer.Status()
	actualHttpBody := response.Body.String()

	assert.Equal(t, expectedHttpStatus, actualHttpStatus)
	assert.Equal(t, expectedHttpBody, actualHttpBody)
	assert.Contains(t, actualHttpBody, expectedError)
}
func Test_GetByID_InternalErrorGetById_500InternalServerError(t *testing.T) {
	expectedHttpStatus := http.StatusInternalServerError
	expectedError := "Error to get user"
	expectedHttpBody := "{\"error\":\"" + expectedError + "\"}"

	userServiceMock := userServiceMock{}
	userServiceMock.getByIDFn = func() (*models.User, error) {
		return nil, errors.New("error stub")
	}
	services.UserService = userServiceMock

	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.AddParam("id", "0")

	UserController.GetUser(context)

	actualHttpStatus := context.Writer.Status()
	actualHttpBody := response.Body.String()

	assert.Equal(t, expectedHttpStatus, actualHttpStatus)
	assert.Equal(t, expectedHttpBody, actualHttpBody)
	assert.Contains(t, actualHttpBody, expectedError)
}

func Test_GetByID_NotFoundId_404NotFound(t *testing.T) {

	expectedHttpStatus := http.StatusNotFound
	expectedError := "No user found"
	expectedHttpBody := "{\"error\":\"" + expectedError + "\"}"

	var userStub models.User
	userStub.ID = 0
	userServiceMock := userServiceMock{}
	userServiceMock.getByIDFn = func() (*models.User, error) {
		return &userStub, nil
	}
	services.UserService = userServiceMock

	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.AddParam("id", "0")

	UserController.GetUser(context)

	actualHttpStatus := context.Writer.Status()
	actualHttpBody := response.Body.String()
	assert.Contains(t, actualHttpBody, expectedError)

	assert.Equal(t, expectedHttpStatus, actualHttpStatus)
	assert.Equal(t, expectedHttpBody, actualHttpBody)
}

func Test_GetByID_FoundId_200ReturnUser(t *testing.T) {

	expectedHttpStatus := http.StatusOK
	expectedHttpBody := "{\"ID\":1,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"DeletedAt\":null,\"name\":\"Kurt\",\"username\":\"user1234\",\"email\":\"\",\"password\":\"encrypted\",\"Events\":null}"

	var userStub models.User
	userStub.ID = 1
	userStub.Name = "Kurt"
	userStub.Username = "user1234"
	userStub.Password = "encrypted"

	userServiceMock := userServiceMock{}
	userServiceMock.getByIDFn = func() (*models.User, error) {
		return &userStub, nil
	}
	services.UserService = userServiceMock

	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.AddParam("id", "1")

	UserController.GetUser(context)

	actualHttpStatus := context.Writer.Status()
	actualHttpBody := response.Body.String()

	assert.Equal(t, expectedHttpStatus, actualHttpStatus)
	assert.Equal(t, expectedHttpBody, actualHttpBody)
}
