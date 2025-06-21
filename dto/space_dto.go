package dto

import (
	"github.com/google/uuid"
)

type CreateSpaceRequest struct {
	Name        string `validate:"required,min=1,max=255" json:"name"`
	Description string `validate:"omitempty,max=1000" json:"description"`
}

type UpdateSpaceRequest struct {
	Id          uuid.UUID `validate:"required,uuid" json:"id"`
	Name        string    `validate:"required,min=1,max=255" json:"name"`
	Description string    `validate:"omitempty,max=1000" json:"description"`
}

type SpaceResponse struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Slug        string    `json:"slug"`
	OwnerID     uuid.UUID `json:"owner_id"`
	MemberCount int       `json:"member_count"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
}
