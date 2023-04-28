package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pavva91/gin-gorm-rest/services"
	"net/http"
)

var (
	PingController = pingController{}
)

type pingController struct{}

func (controller pingController) Ping(context *gin.Context) {
	result, err := services.PingService.HandlePing()
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": result})

	// context.String(http.StatusOK, result)
}
