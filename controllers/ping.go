package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/pavva91/gin-gorm-rest/services"
)

var (
	PingController = pingController{}
)

type pingController struct{}

func (controller pingController) Ping(c *gin.Context) {
	result, err := services.PingService.HandlePing()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusOK, result)
}
