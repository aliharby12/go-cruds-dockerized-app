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
// @Summary Create a new post
// @Description Create a new post with the provided data
// @Accept json
// @Produce json
// @Param post body schema.CreatePostRequest true "Post data in JSON format"
// @Success 201 {object} schema.ViewPostResponse
// @Router /posts [post]
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

// List All Posts
// @Summary List all posts
// @Description Get a list of all posts
// @Produce json
// @Success 200 {object} schema.ListPostsResponse
// @Router /posts [get]
func ListPosts(c *gin.Context) {
	var posts []models.Post
	inits.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{"posts": posts, "count": len(posts), "errors": c.Errors})
}

// View single post
// @Summary View single post
// @Description Get details of a single post by ID
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} schema.ViewPostResponse
// @Router /posts/{id} [get]
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
// @Summary Update post
// @Description Update an existing post
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Param post body schema.UpdatePostRequest true "Updated post data"
// @Success 200 {object} schema.ViewPostResponse
// @Router /posts/{id} [put]
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
// @Summary Delete post
// @Description Delete a post by ID
// @Param id path string true "Post ID"
// @Success 204
// @Router /posts/{id} [delete]
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
