package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"crud-app/controllers"
	"crud-app/inits"
	"crud-app/models"
	"crud-app/utils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreatePost_ValidInput(t *testing.T) {
	// Set the Gin mode to TestMode
	gin.SetMode(gin.TestMode)
	user := setupUser() // Get the created user

	// Create a request to test the CreatePost function
	reqBody := `{"title": "Test Title", "description": "Test Description"}`
	req, _ := http.NewRequest("POST", "/posts", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	token, _ := utils.GenerateToken(user.Username, user.Password, user.ID)
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Mock the Gin context to set the "userID" key
	c.Set("userID", user.ID)

	// Call the CreatePost function
	controllers.CreatePost(c)

	// Assertions
	assert.Equal(t, http.StatusCreated, 201)
}

func TestCreatePost_InvalidInput(t *testing.T) {
	// Set the Gin mode to TestMode
	gin.SetMode(gin.TestMode)
	user := setupUser() // Get the created user

	// Create a request with invalid input
	reqBody := `{"title": "", "description": "Test Description"}`
	req, _ := http.NewRequest("POST", "/posts", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	token, _ := utils.GenerateToken(user.Username, user.Password, user.ID)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Call the CreatePost function
	controllers.CreatePost(c)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, 400)
}

func TestCreatePostWithoutUser(t *testing.T) {
	// Set the Gin mode to TestMode
	gin.SetMode(gin.TestMode)

	// Create a request to test the CreatePost function without a user
	reqBody := `{"title": "Test Title", "description": "Test Description"}`
	req, _ := http.NewRequest("POST", "/posts", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Call the CreatePost function
	controllers.CreatePost(c)

	// Assertions
	assert.Equal(t, http.StatusUnauthorized, 401)
}

func TestListMyPosts_WithPosts(t *testing.T) {
	// Set the Gin mode to TestMode
	gin.SetMode(gin.TestMode)
	user := setupUser()

	// Create a request to test the ListMyPosts function
	req, _ := http.NewRequest("GET", "/posts", nil)
	token, _ := utils.GenerateToken(user.Username, user.Password, user.ID)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Set("userID", user.ID)

	// Call the ListMyPosts function
	controllers.ListMyPosts(c)

	// Assertions
	assert.Equal(t, http.StatusOK, 200)
}

func TestListMyPosts_NoToken(t *testing.T) {
	// Set the Gin mode to TestMode
	gin.SetMode(gin.TestMode)

	// Create a request to test the ListMyPosts function with no token
	req, _ := http.NewRequest("GET", "/my-posts", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Call the ListMyPosts function
	controllers.ListMyPosts(c)

	// Assertions
	assert.Equal(t, http.StatusUnauthorized, 401)
}

func teardownUser() {
	// Delete the user created during setup.
	result := inits.DB.Delete(&models.User{}, createdUserID)
	if result.Error != nil {
		panic("Failed to delete the user after testing: " + result.Error.Error())
	}
}
