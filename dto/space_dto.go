package dto

import (
	"github.com/google/uuid"
)

type CreateSpaceRequest struct {
	Name        string `validate:"required,min=1,max=255" json:"name"`
	Description string `validate:"omitempty,max=1000" json:"description"`
	OwnerId     string `validate:"required,uuid" json:"owner_id"`
}

type UpdateSpaceRequest struct {
	Id          uuid.UUID `validate:"required,uuid" json:"id"`
	Name        string    `validate:"required,min=1,max=255" json:"name"`
	Description string    `validate:"omitempty,max=1000" json:"description"`
}
