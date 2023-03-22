package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pavva91/gin-gorm-rest/config"
	"github.com/pavva91/gin-gorm-rest/models"
)

func ListEvents(c *gin.Context) {
	events := []models.Event{}
	config.DB.Find(&events)
	c.JSON(200, &events)
}

func GetEvent(c *gin.Context) {
	var event models.Event
	config.DB.Where("id = ?", c.Param("id")).First(&event)
	c.JSON(200, &event)
}

func CreateEvent(c *gin.Context) {
	var event models.Event
	c.BindJSON(&event)
	config.DB.Create(&event)
	c.JSON(200, &event)
}

func DeleteEvent(c *gin.Context) {
	var event models.Event
	config.DB.Where("id = ?", c.Param("id")).Delete(&event)
	c.JSON(200, &event)
}

func UpdateEvent(c *gin.Context) {
	var event models.Event
	config.DB.Where("id = ?", c.Param("id")).First(&event)
	c.BindJSON(&event)
	config.DB.Save(&event)
	c.JSON(200, &event)
}
