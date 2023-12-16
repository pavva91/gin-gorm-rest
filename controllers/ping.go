package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pavva91/gin-gorm-rest/services"
	"net/http"
)

var (
	Ping = &ping{}
)

type ping struct{}

func (controller ping) Ping(c *gin.Context) {
	result, err := services.Ping.HandlePing()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": result})

	// c.String(http.StatusOK, result)
}
