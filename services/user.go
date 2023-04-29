package services

import (
	"github.com/pavva91/gin-gorm-rest/db"
	"github.com/pavva91/gin-gorm-rest/models"
)

var (
	UserService userService = userServiceImpl{}
)

type userService interface {
	ListUsers() ([]models.User, error)
	GetByID(id string) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(id string) (*models.User, error)
}

type userServiceImpl struct{}

func (service userServiceImpl) ListUsers() ([]models.User, error) {
	users := []models.User{}
	db.GetDB().Find(&users)
	return users, nil
}

func (service userServiceImpl) GetByID(id string) (*models.User, error) {
	var user *models.User
	db.GetDB().Where("id = ?", id).First(&user)
	return user, nil
}

func (service userServiceImpl) GetByUsername(username string) (*models.User, error) {
	var user *models.User
	db.GetDB().Where("username = ?", username).First(&user)
	return user, nil
}

func (service userServiceImpl) Update(user *models.User) (*models.User, error) {
	db.GetDB().Save(&user)
	return user, nil
}

func (service userServiceImpl) Delete(id string) (*models.User, error) {
	var user *models.User
	db.GetDB().Where("id = ?", id).Delete(&user)
	return user, nil
}

