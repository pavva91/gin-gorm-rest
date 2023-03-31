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
		api.POST("/token", controllers.GenerateToken)
		secured := api.Group("secured").Use(middlewares.Auth())
		{
		}

		healthGroup := api.Group("health")
		{
			healthController := new(controllers.HealthController)
			healthGroup.GET("", healthController.Status)
			// api.GET("/ping", healthController.Ping)
			secured.GET("/ping", healthController.Ping)
		}

		usersGroup := api.Group("users")
		{
			users := new(controllers.UserController)
			usersGroup.GET("/:id", users.Retrieve)
			usersGroup.POST("", controllers.RegisterUser)
		}
		eventsGroup := api.Group("events")
		{
			eventsController := new(controllers.EventController)
			eventsGroup.GET("", eventsController.ListEvents)
			eventsGroup.GET("/", eventsController.ListEvents)
			eventsGroup.GET("/:id", controllers.GetEvent)
			eventsGroup.POST("/", controllers.CreateEvent)
			eventsGroup.DELETE("/:id", controllers.DeleteEvent)
			eventsGroup.PUT("/:id", controllers.SubstituteEvent)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}
