package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Username   string    `gorm:"type:varchar(255);unique;not null" json:"username"`
	Name       string    `gorm:"type:varchar(255)" json:"name"`
	Email      string    `gorm:"type:varchar(255);unique" json:"email"`
	AvatarType string    `gorm:"type:varchar(255)" json:"avatar_type"`
	Password   string    `gorm:"type:varchar(255)" json:"password"`
	CreatedAt  time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:timestamp" json:"updated_at"`
}