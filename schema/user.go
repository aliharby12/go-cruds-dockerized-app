package schema

import "time"

// Swagger model for creating a user
type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ViewUserResponse struct {
	ID        uint       `gorm:"primaryKey" json:"ID"`
	CreatedAt time.Time  `json:"CreatedAt"`
	UpdatedAt time.Time  `json:"UpdatedAt"`
	DeletedAt *time.Time `gorm:"index" json:"DeletedAt,omitempty"`
	Username  string     `json:"Username"`
}
