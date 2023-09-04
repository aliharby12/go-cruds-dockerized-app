package routers

import (
	"crud-app/controllers"

	"github.com/gin-gonic/gin"
)

func InitializeUserRoutes(userGroup *gin.RouterGroup) {

	// Public routes (no authentication required)
	userGroup.POST("/register", controllers.RegisterController)
	userGroup.POST("/login", controllers.LoginController)
}
