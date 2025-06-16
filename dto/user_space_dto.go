package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserSpaceRequest struct {
	UserID  uuid.UUID `json:"user_id" validate:"required,uuid"`
	SpaceID uuid.UUID `json:"space_id" validate:"required,uuid"`
}

type UserSpaceResponse struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	SpaceID   uuid.UUID `json:"space_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
