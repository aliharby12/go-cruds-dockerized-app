package main

import (
	"crud-app/controllers"
	"crud-app/inits"

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

// @securityDefinitions.basic  BasicAuth

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
	{
		posts := v1.Group("/posts")
		{
			posts.GET(":id", controllers.ViewPost)
			posts.GET("", controllers.ListPosts)
			posts.POST("", controllers.CreatePost)
			posts.DELETE(":id", controllers.DeletePost)
			posts.PATCH(":id", controllers.UpdatePost)
		}
	}

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
