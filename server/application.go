package server

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pavva91/gin-gorm-rest/config"
	"github.com/pavva91/gin-gorm-rest/db"
	"github.com/pavva91/gin-gorm-rest/docs"
	"github.com/pavva91/gin-gorm-rest/models"
)

func StartApplication() {
	var cfg config.ServerConfig

	// Get Configs
	err := cleanenv.ReadConfig("./config/config.yml", &cfg)
	if err != nil {
		log.Println(err)
	}

	// Set Swagger Info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "Insert here REST API Description"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = fmt.Sprintf("/%s/%s", cfg.Server.ApiPath, cfg.Server.ApiVersion)
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Connect to DB
	db.DbOrm.ConnectToDB(cfg)
	db.DbOrm.GetDB().AutoMigrate(&models.User{}, &models.Event{})

	// Create Router
	router := NewRouter(cfg)

	MapUrls(cfg)

	// Start Server
	router.Run(cfg.Server.Host + ":" + cfg.Server.Port)
}
