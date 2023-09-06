package controllers

import (
	"crud-app/inits"
	"crud-app/models"
	"crud-app/schema"
	"crud-app/utils"
	"crud-app/validators"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// RegisterController handles user registration.
// @Summary Register a new user
// @Description Register a new user and generate a JWT token
// @Tags Users
// @Accept json
// @Produce json
// @Param post body schema.CreateUserRequest true "User data in JSON format"
// @Success 201 {object} schema.TokenResponse
// @Router /users/register [post]
func RegisterController(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the role is specified in the request JSON
	role, ok := c.Get("role")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role not specified"})
		return
	}

	// Hash the password before saving it to the database
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash the password"})
		return
	}

	// Create a new user model with the hashed password
	newUser := models.User{
		Username: user.Username,
		Password: hashedPassword,
		Role:     role.(string),
	}

	// validate the new user username and password
	if err := validators.ValidateUser(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Save the user model to the database using GORM
	if err := inits.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a JWT access token for the newly registered user
	accessToken, err := utils.GenerateToken(newUser.Username, newUser.Role, newUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"access_token": accessToken})
}

// HashPassword securely hashes the user's password before saving it.
func HashPassword(password string) (string, error) {
	// Generate a salted hash for the password.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// LoginController handles user login.
// LoginController handles user login.
// @Summary Login a user
// @Description Login a user and generate a JWT token
// @Tags Users
// @Accept json
// @Produce json
// @Param post body schema.CreateUserRequest true "User data in JSON format"
// @Success 200 {object} schema.TokenResponse
// @Router /users/login [post]
func LoginController(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the user from the database by username
	var storedUser models.User
	if err := inits.DB.Where("username = ?", user.Username).First(&storedUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Compare the stored hashed password with the provided password
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Authentication successful; generate a JWT access token
	accessToken, err := utils.GenerateToken(storedUser.Username, storedUser.Role, storedUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}

// List All Users
// @Summary List all users
// @Description Get a list of all users
// @Produce json
// @Tags Users
// @Success 200 {object} schema.ListUsersResponse
// @Router /users [get]
func ListUsers(c *gin.Context) {
	var users []models.User
	result := inits.DB.Table("users").
		Select("id, created_at, updated_at, deleted_at, username, role").
		Where("deleted_at IS NULL").
		Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	userResponses := make([]schema.ViewUserResponse, len(users))
	for i, user := range users {
		userResponses[i] = schema.ViewUserResponse{
			Username:  user.Username,
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: &user.DeletedAt.Time,
			Role:      user.Role,
		}
	}

	c.JSON(http.StatusOK, gin.H{"users": userResponses, "count": len(users), "errors": result.Error})
}

// View my profile
// @Summary View my profile
// @Description Get details of my profile
// @Produce json
// @Tags Users
// @Success 200 {object} schema.ViewUserResponse
// @Router /users/profile [get]
func GetMyProfile(c *gin.Context) {
	user, err := GetUserByID(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	userResponse := schema.ViewUserResponse{
		Username:  user.Username,
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: &user.DeletedAt.Time,
		Role:      user.Role,
	}
	c.JSON(http.StatusOK, gin.H{"user": userResponse})
}

// List All User Posts
// @Summary List all user posts
// @Description Get a list of all user posts
// @Param id path string true "User ID"
// @Produce json
// @Tags Users
// @Success 200 {object} schema.ListPostsResponse
// @Router /users/user-posts/{id} [get]
func GetUserPosts(c *gin.Context) {
	userID := c.Param("userID")

	var posts []models.Post
	result := inits.DB.Table("posts").
		Select("id, author_id, title, description, created_at, updated_at").
		Where("author_id = ?", userID).
		Find(&posts)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts, "count": len(posts), "error": result.Error})
}

// get user by id
func GetUserByID(c *gin.Context) (models.User, error) {
	var user models.User
	userID, _ := c.Get("userID")
	result := inits.DB.Table("users").
		Select("id, created_at, updated_at, deleted_at, username, role").
		Where("id = ?", userID, "deleted_at IS NULL").
		First(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}
