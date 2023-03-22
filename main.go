package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pavva91/gin-gorm-rest/config"
	"github.com/pavva91/gin-gorm-rest/routes"
)

// import "github.com/pavva91/gin-gorm-rest/routes"

func main() {
	// Setup config.yml
	var cfg config.ServerConfig

	err := cleanenv.ReadConfig("./config/config.yml", &cfg)
	if err != nil {
		log.Println(err)
	}

	// https://yewtu.be/watch?v=ZI6HaPKHYsg

	router := gin.Default()

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

	config.ConnectToDB(cfg)

	// add routes
	routes.UserRoute(router)
	routes.EventRoute(router)

	// run router
	// router.Run(":" + cfg.Server.Port)
	router.Run(cfg.Server.Protocol + "://" + cfg.Server.Host + ":" + cfg.Server.Port)
}
