package server

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/pavva91/gin-gorm-rest/config"
)

func Init(cfg config.ServerConfig) {
	apiVersion := fmt.Sprintf("/%s/%s", cfg.Server.ApiPath, cfg.Server.ApiVersion)
	router := NewRouter(apiVersion)

	// CORS Configs
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
		log.Println("Incorrect Dev Environment, interrupt execution")
		os.Exit(1)
	}
	if cfg.Server.Environment == "dev" {
		router.Use(cors.Default())
	} else {
		cors_config := cors.DefaultConfig()
		cors_config.AllowOrigins = []string{"http://localhost:5173"}
		router.Use(cors.New(cors_config))
	}

	// router.Run(config.GetString("server.port"))
	router.Run(cfg.Server.Host + ":" + cfg.Server.Port)
}
