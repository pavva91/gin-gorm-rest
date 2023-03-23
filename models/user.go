package models

import (
	"github.com/pavva91/gin-gorm-rest/db"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       int    `json:"ID" gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h User) GetByID(id string) (*User, error) {
	var user *User
	db.GetDB().Where("id = ?", id).First(&user)
	return user, nil
}
