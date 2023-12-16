package services

import (
	"github.com/pavva91/gin-gorm-rest/models"
	"github.com/pavva91/gin-gorm-rest/repositories"
)

var (
	User userServicer = userServiceImpl{}
)

type userServicer interface {
	Create(u *models.User) (*models.User, error)
	List() ([]models.User, error)
	GetByID(id string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	Update(u *models.User) (*models.User, error)
	Delete(id string) (*models.User, error)
}

type userServiceImpl struct{}

func (s userServiceImpl) Create(u *models.User) (*models.User, error) {
	return repositories.UserRepository.CreateUser(u)
}

func (s userServiceImpl) List() ([]models.User, error) {
	return repositories.UserRepository.ListUsers()
}

func (s userServiceImpl) GetByID(id string) (*models.User, error) {
	return repositories.UserRepository.GetByID(id)
}

func (s userServiceImpl) GetByEmail(email string) (*models.User, error) {
	return repositories.UserRepository.GetByEmail(email)
}

func (s userServiceImpl) GetByUsername(username string) (*models.User, error) {
	return repositories.UserRepository.GetByUsername(username)
}

func (s userServiceImpl) Update(u *models.User) (*models.User, error) {
	return repositories.UserRepository.Update(u)
}

func (s userServiceImpl) Delete(id string) (*models.User, error) {
	return repositories.UserRepository.Delete(id)
}
