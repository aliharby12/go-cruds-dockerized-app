package schema

import "time"

// Swagger model for creating a user
type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Swagger model to view a user
type ViewUserResponse struct {
	ID        uint       `json:"ID"`
	CreatedAt time.Time  `json:"CreatedAt"`
	UpdatedAt time.Time  `json:"UpdatedAt"`
	DeletedAt *time.Time `json:"DeletedAt"`
	Username  string     `json:"Username"`
	Role      string     `json:"Role"`
}

// Swagger model for listing users
type ListUsersResponse struct {
	Users  []ViewUserResponse `json:"users"`
	Count  int                `json:"count"`
	Errors string             `json:"errors"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}
