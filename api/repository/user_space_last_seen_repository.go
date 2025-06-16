package repository

import (
	"errors"
	"go_confess_space-project/model"

	"gorm.io/gorm"
)

type UserSpaceLastSeenRepository interface {
	GetLastSeenByUserAndSpace(userID string, spaceID string) (model.UserSpaceLastSeen, error)
	CreateOrUpdateLastSeen(data model.UserSpaceLastSeen) (model.UserSpaceLastSeen, error)
	DeleteLastSeenByUserAndSpace(userID string, spaceID string) error
}

type UserSpaceLastSeenRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserSpaceLastSeenRepositoryImpl(Db *gorm.DB) UserSpaceLastSeenRepository {
	return &UserSpaceLastSeenRepositoryImpl{Db: Db}
}

func (r *UserSpaceLastSeenRepositoryImpl) GetLastSeenByUserAndSpace(userID string, spaceID string) (model.UserSpaceLastSeen, error) {
	var lastSeen model.UserSpaceLastSeen
	result := r.Db.Where("user_id = ? AND space_id = ?", userID, spaceID).First(&lastSeen)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return model.UserSpaceLastSeen{}, nil
	} else if result.Error != nil {
		return model.UserSpaceLastSeen{}, result.Error
	}

	return lastSeen, nil
}

func (r *UserSpaceLastSeenRepositoryImpl) CreateOrUpdateLastSeen(data model.UserSpaceLastSeen) (model.UserSpaceLastSeen, error) {
	var lastSeenRecord model.UserSpaceLastSeen
	result := r.Db.Where("user_id = ? AND space_id = ?", data.UserID, data.SpaceID).First(&lastSeenRecord)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		if err := r.Db.Create(&data).Error; err != nil {
			return model.UserSpaceLastSeen{}, err
		}
		return data, nil
	} else if result.Error != nil {
		return model.UserSpaceLastSeen{}, result.Error
	}

	lastSeenRecord.LastSeen = data.LastSeen
	if err := r.Db.Save(&lastSeenRecord).Error; err != nil {
		return model.UserSpaceLastSeen{}, err
	}
	return lastSeenRecord, nil
}

func (r *UserSpaceLastSeenRepositoryImpl) DeleteLastSeenByUserAndSpace(userID string, spaceID string) error {
	result := r.Db.Where("user_id = ? AND space_id = ?", userID, spaceID).Delete(&model.UserSpaceLastSeen{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
