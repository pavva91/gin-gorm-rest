package repositories

import (
	"github.com/pavva91/gin-gorm-rest/db"
	"github.com/pavva91/gin-gorm-rest/models"
)

var (
	UserRepository userRepository = userRepositoryImpl{}
)

type userRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	ListUsers() ([]models.User, error)
	GetByID(id string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(id string) (*models.User, error)
}

type userRepositoryImpl struct{}

func (repository userRepositoryImpl) CreateUser(user *models.User) (*models.User, error) {
	err := db.DbOrm.GetDB().Create(&user).Error
	return user, err
}

func (repository userRepositoryImpl) ListUsers() ([]models.User, error) {
	users := []models.User{}
	db.DbOrm.GetDB().Find(&users)
	return users, nil
}

func (repository userRepositoryImpl) GetByID(id string) (*models.User, error) {
	var user *models.User
	db.DbOrm.GetDB().Where("id = ?", id).First(&user)
	return user, nil
}

func (repository userRepositoryImpl) GetByEmail(email string) (*models.User, error) {
	var user *models.User
	db.DbOrm.GetDB().Where("email = ?", email).First(&user)
	return user, nil
}

func (repository userRepositoryImpl) GetByUsername(username string) (*models.User, error) {
	var user *models.User
	db.DbOrm.GetDB().Where("username = ?", username).First(&user)
	return user, nil
}

func (repository userRepositoryImpl) Update(user *models.User) (*models.User, error) {
	db.DbOrm.GetDB().Save(&user)
	return user, nil
}

func (repository userRepositoryImpl) Delete(id string) (*models.User, error) {
	var user *models.User
	db.DbOrm.GetDB().Where("id = ?", id).Delete(&user)
	return user, nil
}
