package dto

import "github.com/google/uuid"

type MessageRequest struct {
	SpaceID uuid.UUID `json:"space_id" validate:"required,uuid"`
	UserID  uuid.UUID `json:"user_id" validate:"required,uuid"`
	Content string    `json:"content" validate:"required"`
}
