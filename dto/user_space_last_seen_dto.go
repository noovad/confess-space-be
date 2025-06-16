package dto

import "time"

type UserSpaceLastSeenRequest struct {
	UserID   string    `json:"user_id" binding:"required"`
	SpaceID  string    `json:"space_id" binding:"required"`
	LastSeen time.Time `json:"last_seen" binding:"required"`
}
