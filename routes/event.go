package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pavva91/gin-gorm-rest/controllers"
)

func EventRoute(router *gin.Engine) {
	eventPath := "/events"
	router.GET(eventPath, controllers.ListEvents)
	router.GET(eventPath+"/", controllers.ListEvents)
	router.GET(eventPath+"/:id", controllers.GetEvent)
	router.POST(eventPath+"/", controllers.CreateEvent)
	router.DELETE(eventPath+"/:id", controllers.DeleteEvent)
	router.PUT(eventPath+"/:id", controllers.UpdateEvent)
}

func enableCors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
}

