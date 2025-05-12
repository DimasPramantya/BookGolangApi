// @title       Bookstore API
// @version     1.0
// @description API for managing books and categories with authentication.\n\nTo authorize, click "Authorize" and enter your JWT token in this format:\n**Bearer &lt;your_token&gt;**
// @BasePath    /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"github.com/DimasPramantya/goMiniProject/internal/api/routers"
	"github.com/DimasPramantya/goMiniProject/internal/configs"
	"github.com/DimasPramantya/goMiniProject/internal/database/connection"
	"github.com/DimasPramantya/goMiniProject/internal/database/migration"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/gin-gonic/gin"
	_ "github.com/DimasPramantya/goMiniProject/docs"
)

func main() {
	configs.Initiator()
	connection.Initiator()
	migration.Initiator(connection.DBConnections)
	defer connection.DBConnections.Close()

	r := gin.Default()

	routers.Init(r, connection.DBConnections)

	r.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}