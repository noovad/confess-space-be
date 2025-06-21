package repository

import (
	"errors"
	"go_confess_space-project/dto"
	"go_confess_space-project/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SpaceRepository interface {
	CreateSpace(space model.Space) (model.Space, error)
	GetOwnSpace(ownerID uuid.UUID) (model.Space, error)
	GetSpaces(limit int, page int, search string, isSuggest bool, userId string) (dto.SpaceListResponse, error)
	GetSpaceBySlug(slug string) (dto.SpaceResponse, error)
	UpdateSpace(id uuid.UUID, space model.Space) (model.Space, error)
	DeleteSpace(id uuid.UUID) error
	ExistsByOwnerID(ownerID uuid.UUID) (bool, error)
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

func (t *SpaceRepositoryImpl) GetOwnSpace(ownerID uuid.UUID) (model.Space, error) {
	var space model.Space
	result := t.Db.Where("owner_id = ?", ownerID).First(&space)
	if result.Error != nil {
		return model.Space{}, result.Error
	}
	return space, nil
}

func (t *SpaceRepositoryImpl) GetSpaces(limit int, page int, search string, isSuggest bool, userId string) (dto.SpaceListResponse, error) {
	var spaces []dto.SpaceResponse
	var userSpaces []model.UserSpace
	var query *gorm.DB

	if isSuggest {
		t.Db.Model(&model.UserSpace{}).Where("user_id = ?", userId).Select("space_id").Scan(&userSpaces)
		var excludedIDs []uuid.UUID
		for _, us := range userSpaces {
			excludedIDs = append(excludedIDs, us.SpaceID)
		}
		query = t.Db.Table("spaces s").
			Select(`s.id, s.name, s.slug, s.description, s.owner_id, s.created_at, s.updated_at, COUNT(us.user_id) as member_count`).
			Joins(`LEFT JOIN user_spaces us ON us.space_id = s.id`)
		if len(excludedIDs) > 0 {
			query = query.Where("s.id NOT IN ?", excludedIDs)
		}
	} else {
		query = t.Db.Table("spaces s").
			Select(`s.id, s.name, s.slug, s.description, s.owner_id, s.created_at, s.updated_at, COUNT(us.user_id) as member_count`).
			Joins(`LEFT JOIN user_spaces us ON us.space_id = s.id`)
	}

	if search != "" {
		query = query.Where("s.name ILIKE ?", "%"+search+"%")
	}

	offset := (page - 1) * limit
	result := query.
		Group("s.id").
		Limit(limit).
		Offset(offset).
		Scan(&spaces)

	if result.Error != nil {
		return dto.SpaceListResponse{}, result.Error
	}

	var total int64
	t.Db.Model(&model.Space{}).Count(&total)

	response := dto.SpaceListResponse{
		Spaces:     spaces,
		Total:      int(total),
		Limit:      limit,
		Page:       page,
		TotalPages: (int(total) + limit - 1) / limit,
	}

	return response, nil
}

func (t *SpaceRepositoryImpl) GetSpaceBySlug(slug string) (dto.SpaceResponse, error) {
	var space dto.SpaceResponse
	result := t.Db.Table("spaces s").
		Select(`s.id, s.name, s.slug, s.description, s.owner_id, s.created_at, s.updated_at, COUNT(us.user_id) as member_count`).
		Joins(`LEFT JOIN user_spaces us ON us.space_id = s.id`).
		Where("s.slug = ?", slug).
		Group("s.id").
		First(&space)

	if result.Error != nil {
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

func (t *SpaceRepositoryImpl) ExistsByOwnerID(ownerID uuid.UUID) (bool, error) {
	var count int64
	err := t.Db.Model(&model.Space{}).Where("owner_id = ?", ownerID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
