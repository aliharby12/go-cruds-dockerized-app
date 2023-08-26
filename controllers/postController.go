package controllers

import (
	"crud-app/inits"
	"crud-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Function to get a post by ID
func getPostByID(id string) (*models.Post, error) {
	var post models.Post
	if err := inits.DB.First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

// Create a new post
func CreatePost(c *gin.Context) {
	var body models.Post
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if body.Title == "" || body.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title and description cannot be empty"})
		return
	}

	result := inits.DB.Create(&body)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"post": body})
}

// List posts
func ListPosts(c *gin.Context) {
	var posts []models.Post
	inits.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{"posts": posts, "count": len(posts)})
}

// View single post
func ViewPost(c *gin.Context) {
	id := c.Param("id")
	post, err := getPostByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}

// Update post
func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	post, err := getPostByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	var body models.Post
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if body.Title == "" || body.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title and description cannot be empty"})
		return
	}

	inits.DB.Model(&post).Updates(body)

	c.JSON(http.StatusOK, gin.H{"post": post})
}

// Delete post
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	post, err := getPostByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	if err := inits.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"post": post})
}
