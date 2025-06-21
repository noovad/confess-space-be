package repository

import (
	"go_confess_space-project/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserSpaceRepository interface {
	AddUserToSpace(userSpace model.UserSpace) (model.UserSpace, error)
	RemoveUserFromSpace(spacID, userID uuid.UUID) error
	GetUserSpace(spaceID, userID uuid.UUID) ([]model.UserSpace, error)
	IsUserInSpace(spaceID, userID uuid.UUID) (bool, error)
}

func NewUserSpaceRepositoryImpl(Db *gorm.DB) UserSpaceRepository {
	return &UserSpaceRepositoryImpl{Db: Db}
}

type UserSpaceRepositoryImpl struct {
	Db *gorm.DB
}

func (r *UserSpaceRepositoryImpl) AddUserToSpace(userSpace model.UserSpace) (model.UserSpace, error) {
	result := r.Db.Create(&userSpace)

	if result.Error != nil {
		return userSpace, result.Error
	}

	return userSpace, nil
}

func (r *UserSpaceRepositoryImpl) RemoveUserFromSpace(spaceID uuid.UUID, userID uuid.UUID) error {
	result := r.Db.Where("space_ID = ? AND user_ID = ?", spaceID, userID).Delete(&model.UserSpace{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *UserSpaceRepositoryImpl) GetUserSpace(spaceID uuid.UUID, userID uuid.UUID) ([]model.UserSpace, error) {
	var userSpaces []model.UserSpace
	query := r.Db.Model(&model.UserSpace{}).Preload("User").Preload("Space")

	if spaceID != uuid.Nil {
		query = query.Where("space_ID = ?", spaceID)
	}

	if userID != uuid.Nil {
		query = query.Where("user_ID = ?", userID)
	}

	result := query.Find(&userSpaces)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return []model.UserSpace{}, nil
		}
		return nil, result.Error
	}

	return userSpaces, nil
}

func (r *UserSpaceRepositoryImpl) IsUserInSpace(spaceID, userID uuid.UUID) (bool, error) {
	result := r.Db.Where("space_ID = ? AND  user_ID = ?", spaceID, userID).First(&model.UserSpace{})
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, result.Error
	}

	if result.RowsAffected > 0 {
		return true, nil
	}

	return false, nil
}
