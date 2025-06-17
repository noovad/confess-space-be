package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	Id         uuid.UUID `json:"id"`
	Username   string    `json:"username"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	AvatarType string    `json:"avatar_type"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
