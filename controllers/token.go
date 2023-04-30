package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pavva91/gin-gorm-rest/auth"
	"github.com/pavva91/gin-gorm-rest/errorhandling"
	"github.com/pavva91/gin-gorm-rest/services"
	"github.com/rs/zerolog/log"
)

type TokenRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var (
	TokenController = tokenController{}
)

type tokenController struct{}

func (controller tokenController) GenerateToken(c *gin.Context) {
	var request TokenRequest
	// var user models.User
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
	user, err := services.UserService.GetByEmail(request.Email)
	if err != nil {
		errorMessage := errorhandling.SimpleErrorMessage{Message: err.Error()}
		c.JSON(http.StatusInternalServerError, errorMessage)
		c.Abort()
		return
	}
	if user.ID == 0 {
		errorMessage := errorhandling.SimpleErrorMessage{Message: "User not found"}
		c.JSON(http.StatusUnauthorized, errorMessage)
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
	// tokenString, err := auth.GenerateJWT(user.Email, user.Username)
	tokenString, err := auth.AuthenticationUtility.GenerateJWT(user.Email, user.Username)
	if err != nil {
		errorMessage := errorhandling.SimpleErrorMessage{Message: err.Error()}
		c.JSON(http.StatusInternalServerError, errorMessage)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
