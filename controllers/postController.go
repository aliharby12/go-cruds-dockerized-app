package controllers

import (
	"crud-app/inits"
	"crud-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {

	var body struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Description: body.Description}

	result := inits.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}