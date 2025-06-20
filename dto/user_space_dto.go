package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserSpaceRequest struct {
	UserId  string `json:"user_id" validate:"required"`
	SpaceId string `json:"space_id" validate:"required"`
}

type UserSpaceResponse struct {
	Id        uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"user_id"`
	SpaceId   uuid.UUID `json:"space_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
