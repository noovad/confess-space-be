package service

import (
	"go_confess_space-project/api/repository"
	"go_confess_space-project/dto"
	"go_confess_space-project/helper"
	customerror "go_confess_space-project/helper/customerrors"
	"go_confess_space-project/model"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type SpaceService interface {
	CreateSpace(user dto.CreateSpaceRequest, id string) (model.Space, error)
	GetOwnSpace(ownerID uuid.UUID) (model.Space, error)
	GetSpaces(limit int, page int, search string, isSuggest bool, userId string) (dto.SpaceListResponse, error)
	GetSpaceBySlug(slug string) (dto.SpaceResponse, error)
	UpdateSpace(requestBody dto.UpdateSpaceRequest) (model.Space, error)
	DeleteSpace(id uuid.UUID) error
	ExistsByOwnerID(ownerID uuid.UUID) (bool, error)
}

func NewSpaceServiceImpl(userRepository repository.SpaceRepository, validate *validator.Validate) SpaceService {
	return &SpaceServiceImpl{
		SpaceRepository: userRepository,
		Validate:        validate,
	}
}

type SpaceServiceImpl struct {
	SpaceRepository repository.SpaceRepository
	Validate        *validator.Validate
}

func (t *SpaceServiceImpl) CreateSpace(req dto.CreateSpaceRequest, id string) (model.Space, error) {
	err := t.Validate.Struct(req)
	if err != nil {
		return model.Space{}, customerror.WrapValidation(err)
	}

	spaceModel := model.Space{
		Name:        req.Name,
		Slug:        helper.ToSlug(req.Name),
		Description: req.Description,
		OwnerID:     uuid.MustParse(id),
	}

	createdSpace, err := t.SpaceRepository.CreateSpace(spaceModel)
	if err != nil {
		return model.Space{}, customerror.HandlePostgresError(err)
	}
	return createdSpace, nil
}

func (t *SpaceServiceImpl) GetOwnSpace(ownerID uuid.UUID) (model.Space, error) {
	space, err := t.SpaceRepository.GetOwnSpace(ownerID)
	if err != nil {
		return model.Space{}, customerror.HandlePostgresError(err)
	}
	return space, nil
}

func (t *SpaceServiceImpl) GetSpaces(limit int, page int, search string, isSuggest bool, userId string) (dto.SpaceListResponse, error) {
	spaces, err := t.SpaceRepository.GetSpaces(limit, page, search, isSuggest, userId)
	if err != nil {
		return dto.SpaceListResponse{}, customerror.HandlePostgresError(err)
	}
	return spaces, nil
}

func (t *SpaceServiceImpl) GetSpaceBySlug(slug string) (dto.SpaceResponse, error) {
	space, err := t.SpaceRepository.GetSpaceBySlug(slug)
	if err != nil {
		return dto.SpaceResponse{}, customerror.HandlePostgresError(err)
	}
	return space, nil
}

func (t *SpaceServiceImpl) UpdateSpace(req dto.UpdateSpaceRequest) (model.Space, error) {
	err := t.Validate.Struct(req)
	if err != nil {
		return model.Space{}, customerror.WrapValidation(err)
	}

	space := model.Space{
		Name:        req.Name,
		Slug:        helper.ToSlug(req.Name),
		Description: req.Description,
	}

	updatedSpace, err := t.SpaceRepository.UpdateSpace(req.Id, space)
	if err != nil {
		return model.Space{}, customerror.HandlePostgresError(err)
	}
	return updatedSpace, nil
}

func (t *SpaceServiceImpl) DeleteSpace(id uuid.UUID) error {
	err := t.SpaceRepository.DeleteSpace(id)
	if err != nil {
		return customerror.HandlePostgresError(err)
	}
	return nil
}

func (t *SpaceServiceImpl) ExistsByOwnerID(ownerID uuid.UUID) (bool, error) {
	exists, err := t.SpaceRepository.ExistsByOwnerID(ownerID)
	if err != nil {
		return false, customerror.HandlePostgresError(err)
	}
	return exists, nil
}
