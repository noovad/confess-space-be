package repository

import (
	"errors"
	"go_confess_space-project/dto"
	"go_confess_space-project/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SpaceRepository interface {
	CreateSpace(space model.Space) (model.Space, error)
	GetSpaces(limit int, page int, search string) ([]model.Space, error)
	GetSpaceById(id uuid.UUID) (model.Space, error)
	UpdateSpace(requestBody dto.UpdateSpaceRequest) (model.Space, error)
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
		return space, result.Error
	}
	return space, nil
}

func (t *SpaceRepositoryImpl) GetSpaces(limit int, page int, search string) ([]model.Space, error) {
	var spaces []model.Space
	query := t.Db.Model(&model.Space{})
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

func (t *SpaceRepositoryImpl) UpdateSpace(requestBody dto.UpdateSpaceRequest) (model.Space, error) {
	var space model.Space

	err := t.Db.First(&space)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return space, gorm.ErrRecordNotFound
	} else if err.Error != nil {
		return space, err.Error
	}

	space.Name = requestBody.Name
	space.Description = requestBody.Description
	space.UpdatedAt = time.Now()

	result := t.Db.Save(&space)
	if result.Error != nil {
		return space, result.Error
	}
	return space, nil
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
