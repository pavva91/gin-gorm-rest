package services

import (
	"github.com/pavva91/gin-gorm-rest/models"
	"github.com/pavva91/gin-gorm-rest/repositories"
)

var (
	UserService userService = userServiceImpl{}
)

type userService interface {
	ListUsers() ([]models.User, error)
	GetByID(id string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(id string) (*models.User, error)
}

type userServiceImpl struct{}

func (service userServiceImpl) ListUsers() ([]models.User, error) {
	return repositories.UserRepository.ListUsers()
}

func (service userServiceImpl) GetByID(id string) (*models.User, error) {
	return repositories.UserRepository.GetByID(id)
}

func (service userServiceImpl) GetByEmail(email string) (*models.User, error) {
	return repositories.UserRepository.GetByEmail(email)
}

func (service userServiceImpl) GetByUsername(username string) (*models.User, error) {
	return repositories.UserRepository.GetByUsername(username)
}

func (service userServiceImpl) Update(user *models.User) (*models.User, error) {
	return repositories.UserRepository.Update(user)
}

func (service userServiceImpl) Delete(id string) (*models.User, error) {
	return repositories.UserRepository.Delete(id)
}
