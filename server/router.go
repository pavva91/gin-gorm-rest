package server

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pavva91/gin-gorm-rest/config"
)

var (
	router = gin.Default()
)

func NewRouter(cfg config.ServerConfig) *gin.Engine {

	// router := gin.Default()
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

	return router
}

func MapUrls(cfg config.ServerConfig) {
	apiVersion := fmt.Sprintf("/%s/%s", cfg.Server.ApiPath, cfg.Server.ApiVersion)
	mapUrls(apiVersion)
}
