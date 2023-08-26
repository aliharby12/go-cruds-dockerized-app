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
	r := gin.Default()
	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts", controllers.ListPosts)
	r.Run()
}
