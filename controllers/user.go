package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pavva91/gin-gorm-rest/auth"
	"github.com/pavva91/gin-gorm-rest/db"
	"github.com/pavva91/gin-gorm-rest/errorhandling"
	"github.com/pavva91/gin-gorm-rest/models"
	"github.com/pavva91/gin-gorm-rest/services"
	"github.com/rs/zerolog/log"
)

var User = &user{}

type user struct{}

func (controller user) RegisterUser(c *gin.Context) {
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
	record := db.DbOrm.GetDB().Create(&user)
	if record.Error != nil {
		errorMessage := errorhandling.SimpleErrorMessage{Message: record.Error.Error()}
		c.JSON(http.StatusInternalServerError, errorMessage)
		c.Abort()
		return
	}
	// c.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
	c.JSON(http.StatusCreated, &user)
}

func (controller user) ListUsers(c *gin.Context) {
	users, err := services.User.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to list users", "error": err})
		c.Abort()
		return
	}
	c.JSON(200, &users)
}

func (controller user) GetUser(c *gin.Context) {
	userId := c.Param("id")
	if !validationUtility.IsEmpty(userId) {
		if !validationUtility.IsInt64(userId) {
			r := errorhandling.SimpleErrorMessage{Message: "Not valid parameter, Insert valid id"}
			c.JSON(http.StatusBadRequest, r)
			c.Abort()
			return
		}
		user, err := services.User.GetByID(userId)
		if err != nil {
			r := errorhandling.SimpleErrorMessage{Message: "Error to get user"}
			c.JSON(http.StatusInternalServerError, r)
			c.Abort()
			return
		}

		if validationUtility.IsZero(int(user.ID)) {
			r := errorhandling.SimpleErrorMessage{Message: "No user found"}
			c.JSON(http.StatusNotFound, r)
			c.Abort()
			return
		} else {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	r := errorhandling.SimpleErrorMessage{Message: "empty id"}
	c.JSON(http.StatusBadRequest, r)
	c.Abort()
	return
}

func (controller user) CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	db.DbOrm.GetDB().Create(&user)
	c.JSON(http.StatusCreated, &user)
}

func (controller user) DeleteUser(c *gin.Context) {
	user, err := services.User.Delete(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to delete user", "error": err})
		c.Abort()
		return
	}
	c.JSON(200, &user)
}

func (controller user) UpdateUser(context *gin.Context) {
	// var newUser models.User
	// Check user type requester
	tokenString := context.GetHeader("Authorization")

	claims, err := auth.DecodeJWT(tokenString)
	if err != nil {
		log.Err(err).Msg("Unable to Decode JWT Token")
	}

	log.Info().Msg("Username: " + claims.Username)
	authenticatedUser, err := services.User.GetByUsername(claims.Username)

	oldUser, err := services.User.GetByID(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error to get user", "error": err})
		context.Abort()
		return
	}

	context.BindJSON(&oldUser)

	if authenticatedUser.ID != oldUser.ID {
		context.JSON(http.StatusMethodNotAllowed, gin.H{"message": "Not allowed", "error": err})
		context.Abort()
		return
	}

	services.User.Update(oldUser)
	// db.GetDB().Save(&user)
	context.JSON(200, &oldUser)
}
