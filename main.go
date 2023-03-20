package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pavva91/gin-gorm-rest/config"
	"github.com/pavva91/gin-gorm-rest/routes"
)

// import "github.com/pavva91/gin-gorm-rest/routes"

func main() {
    // https://yewtu.be/watch?v=ZI6HaPKHYsg

    // router := gin.New()
    router := gin.Default()
    config.Connect()
    routes.UserRoute(router)
    router.Run(":8080")
}
