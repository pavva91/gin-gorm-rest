package server

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pavva91/gin-gorm-rest/config"
	"github.com/pavva91/gin-gorm-rest/controllers"
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

	apiVersionGroup := router.Group(apiVersion)
	{
		healthGroup := apiVersionGroup.Group("health")
		{
			health := new(controllers.HealthController)
			healthGroup.GET("", health.Status)
		}

		usersGroup := apiVersionGroup.Group("users")
		{
			users := new(controllers.UserController)
			usersGroup.GET("/:id", users.Retrieve)
		}
		eventsGroup := apiVersionGroup.Group("events")
		{
			// events := new(controllers.EventController)
			// eventsGroup.GET("/:id", events.Retrieve)
			eventsGroup.GET("", controllers.ListEvents)
			eventsGroup.GET("/", controllers.ListEvents)
			eventsGroup.GET("/:id", controllers.GetEvent)
			eventsGroup.POST("/", controllers.CreateEvent)
			eventsGroup.DELETE("/:id", controllers.DeleteEvent)
			eventsGroup.PUT("/:id", controllers.SubstituteEvent)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}
