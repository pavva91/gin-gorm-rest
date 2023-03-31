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
	"github.com/rs/zerolog/log"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateToken(c *gin.Context) {
	var request TokenRequest
	var user models.User
	if err := c.ShouldBindJSON(&request); err != nil {
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
	// check if email exists and password is correct
	record := db.GetDB().Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		errorMessage := errorhandling.SimpleErrorMessage{Message: record.Error.Error()}
		c.JSON(http.StatusInternalServerError, errorMessage)
		c.Abort()
		return
	}
	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		errorMessage := errorhandling.SimpleErrorMessage{Message: "Invalid Credentials"}
		c.JSON(http.StatusUnauthorized, errorMessage)
		c.Abort()
		return
	}
	tokenString, err := auth.GenerateJWT(user.Email, user.Username)
	if err != nil {
		errorMessage := errorhandling.SimpleErrorMessage{Message: err.Error()}
		c.JSON(http.StatusInternalServerError, errorMessage)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
