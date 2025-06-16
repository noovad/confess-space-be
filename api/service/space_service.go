package service

import (
	"go_confess_space-project/api/repository"
	"go_confess_space-project/dto"
	"go_confess_space-project/model"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type SpaceService interface {
	CreateSpace(user dto.CreateSpaceRequest) (model.Space, error)
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

func (t *SpaceServiceImpl) CreateSpace(user dto.CreateSpaceRequest) (model.Space, error) {
	err := t.Validate.Struct(user)
	if err != nil {
		return model.Space{}, err
	}

	spaceModel := model.Space{
		Name:        user.Name,
		Description: user.Description,
		OwnerID:     uuid.MustParse(user.OwnerId),
	}

	return t.SpaceRepository.CreateSpace(spaceModel)
}

func (t *SpaceServiceImpl) GetSpaces(limit int, page int, search string) ([]model.Space, error) {
	spaces, err := t.SpaceRepository.GetSpaces(limit, page, search)
	if err != nil {
		return nil, err
	}
	return spaces, nil
}
func (t *SpaceServiceImpl) GetSpaceById(id uuid.UUID) (model.Space, error) {
	space, err := t.SpaceRepository.GetSpaceById(id)
	if err != nil {
		return model.Space{}, err
	}
	return space, nil
}
func (t *SpaceServiceImpl) UpdateSpace(requestBody dto.UpdateSpaceRequest) (model.Space, error) {

	err := t.Validate.Struct(requestBody)
	if err != nil {
		return model.Space{}, err
	}

	updatedSpace, err := t.SpaceRepository.UpdateSpace(requestBody)
	if err != nil {
		return model.Space{}, err
	}
	return updatedSpace, nil
}
func (t *SpaceServiceImpl) DeleteSpace(id uuid.UUID) error {
	err := t.SpaceRepository.DeleteSpace(id)
	if err != nil {
		return err
	}
	return nil
}
