package main

import (
	"fmt"
	"log"

	// "github.com/gin-gonic/gin"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pavva91/gin-gorm-rest/config"
	"github.com/pavva91/gin-gorm-rest/db"
	docs "github.com/pavva91/gin-gorm-rest/docs"
	"github.com/pavva91/gin-gorm-rest/models"
	"github.com/pavva91/gin-gorm-rest/server"
	// swaggerfiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
)

// import "github.com/pavva91/gin-gorm-rest/routes"

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
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

	err := cleanenv.ReadConfig("./config/config.yml", &cfg)
	if err != nil {
		log.Println(err)
	}

	docs.SwaggerInfo.BasePath = fmt.Sprintf("/%s/%s", cfg.Server.ApiPath, cfg.Server.ApiVersion)
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)

	db.ConnectToDB(cfg)
	db.GetDB().AutoMigrate(&models.User{}, &models.Event{})

	server.Init(cfg)
}
