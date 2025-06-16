package service

import (
	"go_confess_space-project/api/repository"
	"go_confess_space-project/dto"
	customerror "go_confess_space-project/helper/customerrors"
	"go_confess_space-project/model"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UserSpaceService interface {
	AddUserToSpace(userSpaceRequest dto.UserSpaceRequest) (model.UserSpace, error)
	RemoveUserFromSpace(spaceID uuid.UUID, userID uuid.UUID) error
	GetUserSpace(spaceID uuid.UUID, userID uuid.UUID) ([]model.UserSpace, error)
	IsUserInSpace(userID, spaceID uuid.UUID) (bool, error)
}

func NewUserSpaceServiceImpl(userSpaceRepository repository.UserSpaceRepository, validate *validator.Validate) UserSpaceService {
	return &UserSpaceServiceImpl{
		UserSpaceRepository: userSpaceRepository,
		Validate:            validate,
	}
}

type UserSpaceServiceImpl struct {
	UserSpaceRepository repository.UserSpaceRepository
	Validate            *validator.Validate
}

func (s *UserSpaceServiceImpl) AddUserToSpace(req dto.UserSpaceRequest) (model.UserSpace, error) {
	err := s.Validate.Struct(req)
	if err != nil {
		return model.UserSpace{}, customerror.WrapValidation(err)
	}

	userSpaceModel := model.UserSpace{
		UserID:  req.UserID,
		SpaceID: req.SpaceID,
	}

	userSpace, err := s.UserSpaceRepository.AddUserToSpace(userSpaceModel)
	if err != nil {
		return model.UserSpace{}, customerror.HandlePostgresError(err)
	}
	return userSpace, nil
}

func (s *UserSpaceServiceImpl) RemoveUserFromSpace(spaceID uuid.UUID, userID uuid.UUID) error {
	return s.UserSpaceRepository.RemoveUserFromSpace(spaceID, userID)
}

func (s *UserSpaceServiceImpl) GetUserSpace(spaceID uuid.UUID, userID uuid.UUID) ([]model.UserSpace, error) {
	return s.UserSpaceRepository.GetUserSpace(spaceID, userID)
}

func (s *UserSpaceServiceImpl) IsUserInSpace(userID, spaceID uuid.UUID) (bool, error) {
	return s.UserSpaceRepository.IsUserInSpace(userID, spaceID)
}
