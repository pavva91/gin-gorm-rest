package server

import (
	"github.com/gin-gonic/gin"
	"github.com/pavva91/gin-gorm-rest/controllers"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(apiVersion string) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	// Add routes
	router.GET("/health", health.Status)
	// TODO: Understand AuthMiddleware
	// router.Use(middlewares.AuthMiddleware())

	v1 := router.Group(apiVersion)
	{
		usersGroup := v1.Group("users")
		{
			users := new(controllers.UserController)
			usersGroup.GET("/:id", users.Retrieve)
		}
		eventsGroup := v1.Group("events")
		{
			// events := new(controllers.EventController)
			// eventsGroup.GET("/:id", events.Retrieve)
			eventsGroup.GET("", controllers.ListEvents)
			eventsGroup.GET("/", controllers.ListEvents)
			eventsGroup.GET("/:id", controllers.GetEvent)
			eventsGroup.POST("/", controllers.CreateEvent)
			eventsGroup.DELETE("/:id", controllers.DeleteEvent)
			eventsGroup.PUT("/:id", controllers.UpdateEvent)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}
