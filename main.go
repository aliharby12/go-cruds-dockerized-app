package main

import (
	"crud-app/inits"
	"net/http"

	"github.com/gin-gonic/gin"
)

// this function runs before any other functions even main
func init() {
	inits.LoadEnvVariables()
	inits.ConnectDB()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
