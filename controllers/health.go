package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthController struct{}

// HealthController godoc
//
//	@Summary		Check Status
//	@Description	Check the status of the REST API
//	@Tags			health
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	message
//	@Router			/health [get]
func (h HealthController) Status(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

func (h HealthController) Ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "pong"})
}
