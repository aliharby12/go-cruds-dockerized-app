package routers

import (
	"crud-app/controllers"
	"crud-app/middlewares"

	"github.com/gin-gonic/gin"
)

func InitializeUserRoutes(userGroup *gin.RouterGroup) {

	// Public routes (no authentication required)
	userGroup.POST("/register", controllers.RegisterController)
	userGroup.POST("/login", controllers.LoginController)

	// private routers
	userGroup.GET("", middlewares.AdminRoleMiddleware(), controllers.ListUsers)
	userGroup.GET("/profile", middlewares.AuthMiddleware(), controllers.GetMyProfile)
	userGroup.GET("/user-posts/:userID", middlewares.AuthMiddleware(), controllers.GetUserPosts)
}
