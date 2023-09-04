package controllers

import (
	"crud-app/inits"
	"crud-app/models"
	"crud-app/utils"
	"crud-app/validators"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// RegisterController handles user registration.
// @Summary Register a new user
// @Description Register a new user and generate a JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param post body schema.CreateUserRequest true "User data in JSON format"
// @Success 201 {object} schema.ViewUserResponse
// @Router /users/register [post]
func RegisterController(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
	accessToken, err := utils.GenerateToken(newUser.Username)
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
// @Tags Authentication
// @Accept json
// @Produce json
// @Param post body schema.CreateUserRequest true "User data in JSON format"
// @Success 200 {object} schema.ViewUserResponse
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
	accessToken, err := utils.GenerateToken(storedUser.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}
