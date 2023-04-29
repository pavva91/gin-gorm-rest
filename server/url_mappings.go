package server

import (
	"github.com/pavva91/gin-gorm-rest/controllers"
	"github.com/pavva91/gin-gorm-rest/middlewares"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func mapUrls(apiVersion string) {
	api := router.Group(apiVersion)
	{
		// unsecured calls
		api.POST("/token", controllers.TokenController.GenerateToken)

		// secured calls
		secured := api.Group("secured").Use(middlewares.Auth())
		{
		}

		healthGroup := api.Group("health")
		{
			healthGroup.GET("", controllers.HealthController.Status)
			// unsecured
			api.GET("/ping", controllers.PingController.Ping)
			// secured
			secured.GET("/ping", controllers.PingController.Ping)
		}

		usersGroup := api.Group("users")
		{
			usersGroup.POST("", controllers.UserController.RegisterUser)
			usersGroup.GET("", controllers.UserController.ListUsers)
			usersGroup.GET("/", controllers.UserController.ListUsers)
			usersGroup.GET("/:id", controllers.UserController.GetUser)
			usersGroup.PATCH("/:id", controllers.UserController.UpdateUser)
		}
		eventsGroup := api.Group("events")
		{
			eventsGroup.GET("", controllers.EventController.ListEvents)
			eventsGroup.GET("/", controllers.EventController.ListEvents)
			eventsGroup.GET("/:id", controllers.EventController.GetEvent)
		}
		securedEventsGroup := eventsGroup.Use(middlewares.Auth())
		{
			securedEventsGroup.POST("/", controllers.EventController.CreateEvent)
			securedEventsGroup.DELETE("/:id", controllers.EventController.DeleteEvent)
			securedEventsGroup.PUT("/:id", controllers.EventController.SubstituteEvent)
		}

	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
