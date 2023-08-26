package main

import (
	"crud-app/controllers"
	"crud-app/inits"

	"github.com/gin-gonic/gin"
)

// this function runs before any other functions even main
func init() {
	inits.LoadEnvVariables()
	inits.ConnectDB()
}

func main() {
	router := gin.Default()
	router.POST("/posts", controllers.CreatePost)
	router.GET("/posts", controllers.ListPosts)
	router.GET("/posts/:id", controllers.ViewPost)
	router.PUT("/posts/:id", controllers.UpdatePost)
	router.DELETE("/posts/:id", controllers.DeletePost)
	router.Run()
}
