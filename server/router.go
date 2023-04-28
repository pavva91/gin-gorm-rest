package server

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pavva91/gin-gorm-rest/config"
	"github.com/pavva91/gin-gorm-rest/controllers"
	"github.com/pavva91/gin-gorm-rest/middlewares"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(cfg config.ServerConfig) *gin.Engine {
	apiVersion := fmt.Sprintf("/%s/%s", cfg.Server.ApiPath, cfg.Server.ApiVersion)

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// CORS Configs based on SERVER_ENVIRONMENT variable
	switch env := cfg.Server.Environment; env {
	case "dev":
		router.Use(cors.Default())
	case "stage":
		log.Println("TODO: Stage environment Setup, for now allow all CORS")
		router.Use(cors.Default())
	case "prod":
		cors_config := cors.DefaultConfig()
		cors_config.AllowOrigins = cfg.Server.CorsAllowedClients
		router.Use(cors.New(cors_config))
	default:
		log.Printf("Incorrect Dev Environment: %s\nInterrupt execution", env)
		os.Exit(1)
	}

	// Add routes
	// TODO: Understand AuthMiddleware
	// router.Use(middlewares.AuthMiddleware())

	api := router.Group(apiVersion)
	{
		// unsecured calls
		api.POST("/token", controllers.GenerateToken)

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
			usersController := new(controllers.UserController)
			usersGroup.POST("", usersController.RegisterUser)
			usersGroup.GET("", usersController.ListUsers)
			usersGroup.GET("/", usersController.ListUsers)
			usersGroup.GET("/:id", usersController.GetUser)
		}
		eventsGroup := api.Group("events")
		{
			eventsGroup.GET("", controllers.EventController.ListE)
			eventsGroup.GET("/", controllers.EventController.ListE)
			eventsGroup.GET("/:id", controllers.GetEvent)
		}
		securedEventsGroup := eventsGroup.Use(middlewares.Auth())
		{
			securedEventsGroup.POST("/", controllers.CreateEvent)
			securedEventsGroup.DELETE("/:id", controllers.DeleteEvent)
			securedEventsGroup.PUT("/:id", controllers.SubstituteEvent)
		}

	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}
