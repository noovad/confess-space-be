package repository

import (
	"errors"
	"go_confess_space-project/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SpaceRepository interface {
	CreateSpace(space model.Space) (model.Space, error)
	GetSpaces(limit int, page int, search string, isSuggest bool, userId string) ([]model.Space, error)
	GetSpaceById(id uuid.UUID) (model.Space, error)
	GetSpaceBySlug(slug string) (model.Space, error)
	UpdateSpace(id uuid.UUID, space model.Space) (model.Space, error)
	DeleteSpace(id uuid.UUID) error
}

func NewSpaceRepositoryImpl(Db *gorm.DB) SpaceRepository {
	return &SpaceRepositoryImpl{Db: Db}
}

type SpaceRepositoryImpl struct {
	Db *gorm.DB
}

func (r *SpaceRepositoryImpl) CreateSpace(space model.Space) (model.Space, error) {
	result := r.Db.Create(&space)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return model.Space{}, gorm.ErrDuplicatedKey
		}
		if errors.Is(result.Error, gorm.ErrForeignKeyViolated) {
			return model.Space{}, gorm.ErrForeignKeyViolated
		}
		return space, result.Error
	}
	return space, nil
}

func (t *SpaceRepositoryImpl) GetSpaces(limit int, page int, search string, isSuggest bool, userId string) ([]model.Space, error) {
	var spaces []model.Space
	var userSpaces []model.UserSpace
	var query *gorm.DB
	if isSuggest {
		t.Db.Model(&model.UserSpace{}).Where("user_id = ?", userId).Select("space_id").Scan(&userSpaces)
		var excludedIDs []uuid.UUID
		for _, us := range userSpaces {
			excludedIDs = append(excludedIDs, us.SpaceID)
		}
		query = t.Db.Model(&model.Space{})
		if len(excludedIDs) > 0 {
			query = query.Where("id NOT IN ?", excludedIDs)
		}
	} else {
		query = t.Db.Model(&model.Space{})
	}
	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}
	offset := (page - 1) * limit
	result := query.Limit(limit).Offset(offset).Find(&spaces)

	if result.Error != nil {
		return nil, result.Error
	}
	return spaces, nil
}

func (t *SpaceRepositoryImpl) GetSpaceById(id uuid.UUID) (model.Space, error) {
	var space model.Space
	result := t.Db.Where(id).First(&space)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return space, gorm.ErrRecordNotFound
	} else if result.Error != nil {
		return space, result.Error
	}
	return space, nil
}

func (t *SpaceRepositoryImpl) GetSpaceBySlug(slug string) (model.Space, error) {
	var space model.Space
	result := t.Db.Where("slug = ?", slug).First(&space)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return space, gorm.ErrRecordNotFound
	} else if result.Error != nil {
		return space, result.Error
	}
	return space, nil
}

func (t *SpaceRepositoryImpl) UpdateSpace(id uuid.UUID, space model.Space) (model.Space, error) {
	result := t.Db.Model(&model.Space{}).Where("id = ?", id).Updates(space)
	if result.Error != nil {
		return space, result.Error
	}

	if result.RowsAffected == 0 {
		return space, gorm.ErrRecordNotFound
	}

	var updatedSpace model.Space
	err := t.Db.Where("id = ?", id).First(&updatedSpace).Error
	if err != nil {
		return space, err
	}

	return updatedSpace, nil
}

func (t *SpaceRepositoryImpl) DeleteSpace(id uuid.UUID) error {
	result := t.Db.Delete(&model.Space{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
