package dto

type MessageRequest struct {
	SpaceID string `json:"space_id" validate:"required"`
	Message string `json:"message" validate:"required"`
}
