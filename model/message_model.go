package model

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	SpaceID   uuid.UUID `gorm:"type:uuid;not null" json:"space_id"`
	Space     *Space    `gorm:"foreignKey:SpaceID;references:ID" json:"space,omitempty"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	User      *User     `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	CreatedAt time.Time `gorm:"type:timestamp;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:current_timestamp" json:"updated_at"`
}
