package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pavva91/gin-gorm-rest/controllers"
)

func UserRoute(router *gin.Engine) {
	userPath := "/users"
	router.GET(userPath, controllers.ListUsers)
	router.GET(userPath+"/", controllers.ListUsers)
	router.GET(userPath+"/:id", controllers.GetUser)
	router.POST(userPath+"/", controllers.CreateUser)
	router.DELETE(userPath+"/:id", controllers.DeleteUser)
	router.PUT(userPath+"/:id", controllers.UpdateUser)
}
