package services

import (
	"errors"
	"testing"

	"github.com/pavva91/gin-gorm-rest/models"
	"github.com/pavva91/gin-gorm-rest/repositories"
	"github.com/pavva91/gin-gorm-rest/stubs"
	"github.com/stretchr/testify/assert"
)

func Test_CreateUser_Error(t *testing.T) {
	// mocks
	user := models.User{}

	// Stubs
	errorMessage := "unexpected error"
	unexpectedError := errors.New(errorMessage)
	userRepositoryStub := stubs.UserRepositoryStub{}
	userRepositoryStub.CreateUserFn = func(*models.User) (*models.User, error){
		return nil, unexpectedError
	}
	repositories.UserRepository = userRepositoryStub

	// Call function to test
	userReturn, err := UserService.CreateUser(&user)

	// Check Values
	assert.NotNil(t, err)
	assert.Equal(t, errorMessage, err.Error())
	assert.Nil(t, userReturn)
}

func Test_CreateUser_OK(t *testing.T) {
	// mocks
	user := models.User{}

	// Stubs
	userRepositoryStub := stubs.UserRepositoryStub{}
	userRepositoryStub.CreateUserFn = func(*models.User) (*models.User, error){
		return &user, nil
	}
	repositories.UserRepository = userRepositoryStub

	// Call function to test
	userReturn, err := UserService.CreateUser(&user)

	// Check Values
	assert.Nil(t, err)
	assert.NotNil(t, userReturn)
	assert.Equal(t, userReturn, &user)
}
