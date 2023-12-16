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
		api.POST("/token", controllers.JWT.GenerateToken)

		// secured calls
		secured := api.Group("secured").Use(middlewares.Auth())
		{
		}

		healthGroup := api.Group("health")
		{
			healthGroup.GET("", controllers.Health.Status)
			// unsecured
			api.GET("/ping", controllers.Ping.Ping)
			// secured
			secured.GET("/ping", controllers.Ping.Ping)
		}

		usersGroup := api.Group("users")
		{
			usersGroup.POST("", controllers.User.RegisterUser)
			usersGroup.GET("", controllers.User.ListUsers)
			usersGroup.GET("/", controllers.User.ListUsers)
			usersGroup.GET("/:id", controllers.User.GetUser)
			usersGroup.PATCH("/:id", controllers.User.UpdateUser)
		}
		eventsGroup := api.Group("events")
		{
			eventsGroup.GET("", controllers.Event.List)
			eventsGroup.GET("/", controllers.Event.List)
			eventsGroup.GET("/:id", controllers.Event.Get)
		}
		securedEventsGroup := eventsGroup.Use(middlewares.Auth())
		{
			securedEventsGroup.POST("/", controllers.Event.Create)
			securedEventsGroup.DELETE("/:id", controllers.Event.DeleteEvent)
			securedEventsGroup.PUT("/:id", controllers.Event.SubstituteEvent)
		}

	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
