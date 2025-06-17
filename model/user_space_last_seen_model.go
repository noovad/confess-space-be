package model

import "time"

type UserSpaceLastSeen struct {
	ID       string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	SpaceID  string    `gorm:"type:uuid;not null" json:"space_id"`
	Space    *Space    `gorm:"foreignKey:SpaceID;references:ID" json:"space,omitempty"`
	UserID   string    `gorm:"type:uuid;not null" json:"user_id"`
	User     *User     `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
	LastSeen time.Time `gorm:"type:timestamp;not null" json:"last_seen"`
}
