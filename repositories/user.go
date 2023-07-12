package repositories

import (
	"github.com/pavva91/gin-gorm-rest/db"
	"github.com/pavva91/gin-gorm-rest/models"
)

var (
	UserRepository userRepository = userRepositoryImpl{}
)

type userRepository interface {
	ListUsers() ([]models.User, error)
	GetByID(id string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(id string) (*models.User, error)
}

type userRepositoryImpl struct{}

func (repository userRepositoryImpl) ListUsers() ([]models.User, error) {
	users := []models.User{}
	db.GetDB().Find(&users)
	return users, nil
}

func (repository userRepositoryImpl) GetByID(id string) (*models.User, error) {
	var user *models.User
	db.GetDB().Where("id = ?", id).First(&user)
	return user, nil
}

func (repository userRepositoryImpl) GetByEmail(email string) (*models.User, error) {
	var user *models.User
	db.GetDB().Where("email = ?", email).First(&user)
	return user, nil
}

func (repository userRepositoryImpl) GetByUsername(username string) (*models.User, error) {
	var user *models.User
	db.GetDB().Where("username = ?", username).First(&user)
	return user, nil
}

func (repository userRepositoryImpl) Update(user *models.User) (*models.User, error) {
	db.GetDB().Save(&user)
	return user, nil
}

func (repository userRepositoryImpl) Delete(id string) (*models.User, error) {
	var user *models.User
	db.GetDB().Where("id = ?", id).Delete(&user)
	return user, nil
}
