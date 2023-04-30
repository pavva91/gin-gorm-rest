package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pavva91/gin-gorm-rest/auth"
	"github.com/pavva91/gin-gorm-rest/mocking"
	"github.com/pavva91/gin-gorm-rest/models"
	"github.com/pavva91/gin-gorm-rest/services"
	"github.com/stretchr/testify/assert"
)

type authenticationUtilityMock struct {
	generateJWTFn func() (tokenString string, err error)
}

func (mock authenticationUtilityMock) GenerateJWT(email string, username string) (tokenString string, err error) {
	return mock.generateJWTFn()
}

func Test_GenerateToken_InvalidRequestBody_400BadRequest(t *testing.T) {
	expectedHttpStatus := http.StatusBadRequest
	expectedError := "Bad Request"
	expectedHttpBody := "{\"error\":\"" + expectedError + "\"}"

	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)

	context.Request = &http.Request{
		Header: make(http.Header),
	}

	wrongRequestBody := []string{"foo", "bar", "baz"}
	mocking.MockJsonPost(context, wrongRequestBody)

	TokenController.GenerateToken(context)

	actualHttpStatus := context.Writer.Status()
	actualHttpBody := response.Body.String()

	assert.Equal(t, expectedHttpStatus, actualHttpStatus)
	assert.Equal(t, expectedHttpBody, actualHttpBody)
}

func Test_GenerateToken_InvalidRequestBodyNoEmailField_400BadRequest(t *testing.T) {
	missingField := "Email"
	expectedHttpStatus := http.StatusBadRequest
	expectedHttpBody := "{\"errors\":[{\"field\":\"" + missingField + "\",\"reason\":\"required\"}]}"

	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)

	context.Request = &http.Request{
		Header: make(http.Header),
	}

	mockRequestBody := TokenRequest{
		Password: "pass1234",
	}

	mocking.MockJsonPost(context, mockRequestBody)

	TokenController.GenerateToken(context)

	actualHttpStatus := context.Writer.Status()
	actualHttpBody := response.Body.String()

	assert.Equal(t, expectedHttpStatus, actualHttpStatus)
	assert.Equal(t, expectedHttpBody, actualHttpBody)
	assert.Contains(t, expectedHttpBody, missingField)
}

func Test_GenerateToken_InvalidRequestBodyNoPasswordField_400BadRequest(t *testing.T) {
	missingField := "Password"
	expectedHttpStatus := http.StatusBadRequest
	expectedHttpBody := "{\"errors\":[{\"field\":\"" + missingField + "\",\"reason\":\"required\"}]}"

	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)

	context.Request = &http.Request{
		Header: make(http.Header),
	}

	mockRequestBody := TokenRequest{
		Email: "alice@wonder.land",
	}

	mocking.MockJsonPost(context, mockRequestBody)

	TokenController.GenerateToken(context)

	actualHttpStatus := context.Writer.Status()
	actualHttpBody := response.Body.String()

	assert.Equal(t, expectedHttpStatus, actualHttpStatus)
	assert.Equal(t, expectedHttpBody, actualHttpBody)
	assert.Contains(t, expectedHttpBody, missingField)
}

func Test_GenerateToken_GetByEmailError_500InternalServerError(t *testing.T) {
	internalErrorMessage := "Internal Error Message"

	expectedHttpStatus := http.StatusInternalServerError
	expectedError := internalErrorMessage
	expectedHttpBody := "{\"error\":\"" + expectedError + "\"}"

	userServiceMock := userServiceMock{}
	userServiceMock.getByEmailFn = func() (*models.User, error) {
		return nil, errors.New(internalErrorMessage)
	}
	services.UserService = userServiceMock

	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)

	context.Request = &http.Request{
		Header: make(http.Header),
	}

	mockRequestBody := TokenRequest{
		Email:    "alice@wonder.land",
		Password: "pass1234",
	}

	mocking.MockJsonPost(context, mockRequestBody)

	TokenController.GenerateToken(context)

	actualHttpStatus := context.Writer.Status()
	actualHttpBody := response.Body.String()

	assert.Equal(t, expectedHttpStatus, actualHttpStatus)
	assert.Equal(t, expectedHttpBody, actualHttpBody)
}

func Test_GenerateToken_WrongEmail_401Unauthorized(t *testing.T) {
	expectedHttpStatus := http.StatusUnauthorized
	expectedError := "User not found"
	expectedHttpBody := "{\"error\":\"" + expectedError + "\"}"

	var userMock models.User
	userMock.ID = 0
	userServiceMock := userServiceMock{}
	userServiceMock.getByEmailFn = func() (*models.User, error) {
		return &userMock, nil
	}
	services.UserService = userServiceMock

	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)

	context.Request = &http.Request{
		Header: make(http.Header),
	}

	mockRequestBody := TokenRequest{

		Email:    "alice@wonder.land",
		Password: "pass1234",
	}

	mocking.MockJsonPost(context, mockRequestBody)

	TokenController.GenerateToken(context)

	actualHttpStatus := context.Writer.Status()
	actualHttpBody := response.Body.String()

	assert.Equal(t, expectedHttpStatus, actualHttpStatus)
	assert.Equal(t, expectedHttpBody, actualHttpBody)
}

func Test_GenerateToken_WrongPassword_401Unauthorized(t *testing.T) {
	correctPassword := "pass1234"
	wrongPassword := "wrong_password"
	expectedHttpStatus := http.StatusUnauthorized
	expectedError := "Invalid Credentials"
	expectedHttpBody := "{\"error\":\"" + expectedError + "\"}"

	var userMock models.User
	userMock.ID = 1
	userMock.HashPassword(correctPassword)
	userServiceMock := userServiceMock{}
	userServiceMock.getByEmailFn = func() (*models.User, error) {
		return &userMock, nil
	}
	services.UserService = userServiceMock

	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)

	context.Request = &http.Request{
		Header: make(http.Header),
	}

	mockRequestBody := TokenRequest{
		Email:    "alice@wonder.land",
		Password: wrongPassword,
	}

	mocking.MockJsonPost(context, mockRequestBody)

	TokenController.GenerateToken(context)

	actualHttpStatus := context.Writer.Status()
	actualHttpBody := response.Body.String()

	assert.Equal(t, expectedHttpStatus, actualHttpStatus)
	assert.Equal(t, expectedHttpBody, actualHttpBody)
}

func Test_GenerateToken_GenerateJWTError_500InternalServerError(t *testing.T) {
	correctPassword := "pass1234"
	internalErrorMessageMock := "Internal Error Message JWT GenerateToken"

	expectedHttpStatus := http.StatusInternalServerError
	expectedError := internalErrorMessageMock
	expectedHttpBody := "{\"error\":\"" + expectedError + "\"}"

	var userMock models.User
	userMock.ID = 1
	userMock.HashPassword(correctPassword)
	userServiceMock := userServiceMock{}
	userServiceMock.getByEmailFn = func() (*models.User, error) {
		return &userMock, nil
	}
	services.UserService = userServiceMock

	authenticationUtilityMock := authenticationUtilityMock{}
	authenticationUtilityMock.generateJWTFn = func() (tokenString string, err error) {
		return "", errors.New(internalErrorMessageMock)
	}
	auth.AuthenticationUtility = authenticationUtilityMock

	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)

	context.Request = &http.Request{
		Header: make(http.Header),
	}

	mockRequestBody := TokenRequest{
		Email:    "alice@wonder.land",
		Password: correctPassword,
	}

	mocking.MockJsonPost(context, mockRequestBody)

	TokenController.GenerateToken(context)

	actualHttpStatus := context.Writer.Status()
	actualHttpBody := response.Body.String()

	assert.Equal(t, expectedHttpStatus, actualHttpStatus)
	assert.Equal(t, expectedHttpBody, actualHttpBody)
}

func Test_GenerateToken_CorrectUser_200JWTToken(t *testing.T) {
	correctPassword := "pass1234"
	tokenStringMock := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFsaWNlODkiLCJlbWFpbCI6ImFsaWNlQGdtYWlsLmNvbSIsImV4cCI6MTY4Mjc4NzgxNiwibmJmIjotNjIxMzU1OTY4MDAsImlhdCI6LTYyMTM1NTk2ODAwfQ.FK1QHfZOs82mZpkzw2PX8E2KfUnDfwrxPjmIpPclVdU"

	expectedHttpStatus := http.StatusOK
	expectedHttpBody := "{\"token\":\"" + tokenStringMock + "\"}"

	var userMock models.User
	userMock.ID = 1
	userMock.HashPassword(correctPassword)
	userServiceMock := userServiceMock{}
	userServiceMock.getByEmailFn = func() (*models.User, error) {
		return &userMock, nil
	}
	services.UserService = userServiceMock

	authenticationUtilityMock := authenticationUtilityMock{}
	authenticationUtilityMock.generateJWTFn = func() (tokenString string, err error) {
		return tokenStringMock, nil
	}
	auth.AuthenticationUtility = authenticationUtilityMock

	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)

	context.Request = &http.Request{
		Header: make(http.Header),
	}

	mockRequestBody := TokenRequest{
		Email:    "alice@wonder.land",
		Password: correctPassword,
	}

	mocking.MockJsonPost(context, mockRequestBody)

	TokenController.GenerateToken(context)

	actualHttpStatus := context.Writer.Status()
	actualHttpBody := response.Body.String()

	assert.Equal(t, expectedHttpStatus, actualHttpStatus)
	assert.Equal(t, expectedHttpBody, actualHttpBody)
}
