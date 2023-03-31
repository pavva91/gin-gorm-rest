package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pavva91/gin-gorm-rest/db"
	"github.com/pavva91/gin-gorm-rest/errorhandling"
	"github.com/pavva91/gin-gorm-rest/models"
	"github.com/rs/zerolog/log"
)

type UserController struct{}

var userModel = new(models.User)

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			errorMessage := errorhandling.ValidationErrorsMessage{Message: errorhandling.NewJSONFormatter().Descriptive(verr)}
			c.JSON(http.StatusBadRequest, errorMessage)
			c.Abort()
			return
		}
		log.Info().Err(err).Msg("unable to bind")
		errorMessage := errorhandling.SimpleErrorMessage{Message: "Bad Request"}
		c.JSON(http.StatusBadRequest, errorMessage)
		c.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		errorMessage := errorhandling.SimpleErrorMessage{Message: err.Error()}
		c.JSON(http.StatusInternalServerError, errorMessage)
		c.Abort()
		return
	}
	record := db.GetDB().Create(&user)
	if record.Error != nil {
		errorMessage := errorhandling.SimpleErrorMessage{Message: record.Error.Error()}
		c.JSON(http.StatusInternalServerError, errorMessage)
		c.Abort()
		return
	}
	// c.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
	c.JSON(http.StatusCreated, &user)
}

func (u UserController) Retrieve(c *gin.Context) {
	if c.Param("id") != "" {
		user, err := userModel.GetByID(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user", "error": err})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User founded!", "user": user})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
	return
}

func ListUsers(c *gin.Context) {
	users := []models.User{}
	db.GetDB().Find(&users)
	c.JSON(200, &users)
}

func GetUser(c *gin.Context) {
	var user models.User
	db.GetDB().Where("id = ?", c.Param("id")).First(&user)
	c.JSON(200, &user)
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	db.GetDB().Create(&user)
	c.JSON(http.StatusCreated, &user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	db.GetDB().Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(200, &user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	db.GetDB().Where("id = ?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	db.GetDB().Save(&user)
	c.JSON(200, &user)
}
