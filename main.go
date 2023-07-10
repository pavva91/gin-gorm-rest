package main

import (
	// docs "github.com/pavva91/gin-gorm-rest/docs"

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

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	server.StartApplication()
}
