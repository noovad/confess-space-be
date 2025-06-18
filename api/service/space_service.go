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
	GetSpaces(limit int, page int, search string) ([]model.Space, error)
	GetSpaceById(id uuid.UUID) (model.Space, error)
	UpdateSpace(requestBody dto.UpdateSpaceRequest) (model.Space, error)
	DeleteSpace(id uuid.UUID) error
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

func (t *SpaceServiceImpl) GetSpaces(limit int, page int, search string) ([]model.Space, error) {
	return t.SpaceRepository.GetSpaces(limit, page, search)
}

func (t *SpaceServiceImpl) GetSpaceById(id uuid.UUID) (model.Space, error) {
	return t.SpaceRepository.GetSpaceById(id)
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

	return t.SpaceRepository.UpdateSpace(req.Id, space)
}

func (t *SpaceServiceImpl) DeleteSpace(id uuid.UUID) error {
	return t.SpaceRepository.DeleteSpace(id)
}
