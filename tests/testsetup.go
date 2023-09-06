package tests

import (
	"crud-app/inits"
	"crud-app/models"

	"github.com/google/uuid"
)

var createdUserID uint

func setupUser() models.User {
	// Create a new user and insert it into the database.
	user := models.User{
		Username: uuid.New().String(),
		Password: "testpassword",
	}

	result := inits.DB.Create(&user)
	if result.Error != nil {
		panic("Failed to create a user for testing: " + result.Error.Error())
	}
	return user
}
