package routers

import (
	"crud-app/controllers"
	"crud-app/middlewares"

	"github.com/gin-gonic/gin"
)

func InitializePostRoutes(postGroup *gin.RouterGroup) {

	postGroup.Use(middlewares.AuthMiddleware())

	postGroup.GET(":id", controllers.ViewPost)
	postGroup.GET("", controllers.ListPosts)
	postGroup.POST("", controllers.CreatePost)
	postGroup.DELETE(":id", controllers.DeletePost)
	postGroup.PATCH(":id", controllers.UpdatePost)
}
