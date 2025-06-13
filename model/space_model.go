package model

import (
	"time"

	"github.com/google/uuid"
)

type Space struct {
	Id          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string    `gorm:"type:varchar(255);unique;not null"`
	Description string    `gorm:"type:text"`
	OwnerId     uuid.UUID `gorm:"type:uuid;not null"`
	Owner       *User     `gorm:"foreignKey:OwnerId;references:Id" json:"owner,omitempty"`
	CreatedAt   time.Time `gorm:"type:timestamp"`
	UpdatedAt   time.Time `gorm:"type:timestamp"`
}
