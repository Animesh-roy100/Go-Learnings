package dto

import (
	"github.com/google/uuid"
)

// CreateUserRequest represents the input for creating a new user
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

// UserResponse represents the response for a user
type UserResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

// VersionResponse represents a version of a user
type VersionResponse struct {
	ID            uuid.UUID              `json:"id"`
	Changes       map[string]interface{} `json:"changes"`
	CommitMessage string                 `json:"commit_message"`
	CreatedAt     string                 `json:"created_at"`
}
