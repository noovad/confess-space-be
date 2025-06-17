package model

import (
	"time"

	"github.com/google/uuid"
)

type UserSpace struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	User      *User     `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
	SpaceID   uuid.UUID `gorm:"type:uuid;not null" json:"space_id"`
	Space     *Space    `gorm:"foreignKey:SpaceID;references:ID" json:"space,omitempty"`
	CreatedAt time.Time `gorm:"type:timestamp"`
	UpdatedAt time.Time `gorm:"type:timestamp"`
}
