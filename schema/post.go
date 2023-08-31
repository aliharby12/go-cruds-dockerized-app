package schema

import "time"

// Swagger model for creating a post
type CreatePostRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Swagger model for listing posts
type ListPostsResponse struct {
	Posts  []ViewPostResponse `json:"posts"`
	Count  int                `json:"count"`
	Errors string             `json:"errors"`
}

// Swagger model for updating a post
type UpdatePostRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Swagger model for viewing a single post
type ViewPostResponse struct {
	ID          uint       `gorm:"primaryKey" json:"ID"`
	CreatedAt   time.Time  `json:"CreatedAt"`
	UpdatedAt   time.Time  `json:"UpdatedAt"`
	DeletedAt   *time.Time `gorm:"index" json:"DeletedAt,omitempty"`
	Title       string     `json:"Title"`
	Description string     `json:"Description"`
}
