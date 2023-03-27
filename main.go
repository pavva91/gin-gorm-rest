package main

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pavva91/gin-gorm-rest/config"
	"github.com/pavva91/gin-gorm-rest/db"
	docs "github.com/pavva91/gin-gorm-rest/docs"
	"github.com/pavva91/gin-gorm-rest/models"
	"github.com/pavva91/gin-gorm-rest/server"
)

// import "github.com/pavva91/gin-gorm-rest/routes"

//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@securityDefinitions.basic	BasicAuth

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
func main() {
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
	db.ConnectToDB(cfg)
	db.GetDB().AutoMigrate(&models.User{}, &models.Event{})

	// Create Router
	router := server.NewRouter(cfg)

	// Start Server
	router.Run(cfg.Server.Host + ":" + cfg.Server.Port)

}
