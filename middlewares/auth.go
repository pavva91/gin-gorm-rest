package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pavva91/gin-gorm-rest/auth"
	"github.com/pavva91/gin-gorm-rest/errorhandling"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			errorMessage := errorhandling.SimpleErrorMessage{Message: "request does not contain an access token"}
			c.JSON(http.StatusBadRequest, errorMessage)
			c.Abort()
			return
		}
		err := auth.ValidateToken(tokenString)
		if err != nil {
			errorMessage := errorhandling.SimpleErrorMessage{Message: err.Error()}
			c.JSON(http.StatusUnauthorized, errorMessage)
			c.Abort()
			return
		}
		c.Next()
	}
}
