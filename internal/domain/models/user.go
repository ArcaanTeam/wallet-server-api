package models

import (
	"time"
)

// User represents the user entity
// @Description User entity with personal information
type User struct {
	ID        string    `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	FirstName string    `json:"first_name" example:"John"`
	LastName  string    `json:"last_name" example:"Doe"`
	Email     string    `json:"email" example:"john.doe@example.com"`
	Age       int       `json:"age" example:"30"`
	CreatedAt time.Time `json:"created_at" example:"2025-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2025-01-01T00:00:00Z"`
}

// UserCreateRequest represents user creation request
// @Description User creation request data
type UserCreateRequest struct {
	FirstName string `json:"first_name" binding:"required,min=1,max=50" example:"John"`
	LastName  string `json:"last_name" binding:"required,min=1,max=50" example:"Doe"`
	Email     string `json:"email" binding:"required,email" example:"john.doe@example.com"`
	Age       int    `json:"age" binding:"required,min=18,max=120" example:"30"`
}

// UserUpdateRequest represents user update request
// @Description user update request data
type UserUpdateRequest struct {
	FirstName *string `json:"first_name,omitempty" example:"john"`
	LastName  *string `json:"last_name,omitempty" example:"john"`
	Age       *int    `json:"age,omitempty" example:"31"`
}
