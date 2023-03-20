package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pavva91/gin-gorm-rest/controllers"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controllers.ListUsers)
	router.GET("/:id", controllers.GetUser)
	router.POST("/", controllers.CreateUser)
	router.DELETE("/:id", controllers.DeleteUser)
	router.PUT("/:id", controllers.UpdateUser)
}
