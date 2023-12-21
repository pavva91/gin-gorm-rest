package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name             string             `gorm:"not null;default:null"`
	Surname          string             `gorm:"not null;default:null"`
	Username         string             `gorm:"unique"`
	Email            string             `gorm:"unique"`
	Password         string             `gorm:"not null;default:null"`
	UsersEvents      []UsersEvents      `gorm:"foreignKey:UserID"`
	OrganizersEvents []OrganizersEvents `gorm:"foreignKey:OrganizerID"`
}

// HashPassword method.
func (user *Users) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	user.Password = string(bytes)

	return nil
}

func (user *Users) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}
