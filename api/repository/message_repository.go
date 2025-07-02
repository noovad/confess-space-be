package repository

import (
	"go_confess_space-project/model"

	"gorm.io/gorm"
)

type MessageRepository interface {
	CreateMessage(model.Message) (model.Message, error)
	GetMessages(spaceID string) ([]model.Message, error)
}

func NewMessageRepositoryImpl(Db *gorm.DB) MessageRepository {
	return &MessageRepositoryImpl{Db: Db}
}

type MessageRepositoryImpl struct {
	Db *gorm.DB
}

func (r *MessageRepositoryImpl) CreateMessage(message model.Message) (model.Message, error) {
	result := r.Db.Create(&message)
	if result.Error != nil {
		return model.Message{}, result.Error
	}

	var createdMessage model.Message
	err := r.Db.
		Preload("User").
		Preload("Space").
		First(&createdMessage, "id = ?", message.ID).Error
	if err != nil {
		return model.Message{}, err
	}

	return createdMessage, nil
}

func (r *MessageRepositoryImpl) GetMessages(spaceID string) ([]model.Message, error) {
	var messages []model.Message
	result := r.Db.
		Where("space_id = ?", spaceID).
		Preload("User").
		Preload("Space").
		Order("created_at DESC").
		Find(&messages)
	if result.Error != nil {
		return nil, result.Error
	}

	return messages, nil
}
