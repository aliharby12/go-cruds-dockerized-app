package main

import (
	"crud-app/inits"
	"crud-app/routers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	"crud-app/docs"
)

// this function runs before any other functions even main
func init() {
	inits.LoadEnvVariables()
	inits.ConnectDB()
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

//@securityDefinitions.apikey Authorization
//@in header
//@name Bearer

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	router := gin.Default()

	// Use CORS middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "CRUDs API"
	docs.SwaggerInfo.Description = "This is a sample CRUDs API server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	v1 := router.Group("/api/v1")

	// Import and initialize user and post routes
	routers.InitializeUserRoutes(v1.Group("/users"))
	routers.InitializePostRoutes(v1.Group("/posts"))

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
