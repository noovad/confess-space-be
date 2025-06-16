package service

import (
	"go_confess_space-project/api/repository"
	"go_confess_space-project/dto"
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

func (s *UserSpaceServiceImpl) AddUserToSpace(userSpaceRequest dto.UserSpaceRequest) (model.UserSpace, error) {
	err := s.Validate.Struct(userSpaceRequest)
	if err != nil {
		return model.UserSpace{}, err
	}

	userSpaceModel := model.UserSpace{
		UserID:  userSpaceRequest.UserID,
		SpaceID: userSpaceRequest.SpaceID,
	}

	userSpace, err := s.UserSpaceRepository.AddUserToSpace(userSpaceModel)
	if err != nil {
		return model.UserSpace{}, err
	}

	return userSpace, nil
}

func (s *UserSpaceServiceImpl) RemoveUserFromSpace(spaceID uuid.UUID, userID uuid.UUID) error {
	err := s.UserSpaceRepository.RemoveUserFromSpace(spaceID, userID)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserSpaceServiceImpl) GetUserSpace(spaceID uuid.UUID, userID uuid.UUID) ([]model.UserSpace, error) {
	userSpaces, err := s.UserSpaceRepository.GetUserSpace(spaceID, userID)
	if err != nil {
		return nil, err
	}

	return userSpaces, nil
}

func (s *UserSpaceServiceImpl) IsUserInSpace(userID, spaceID uuid.UUID) (bool, error) {
	isInSpace, err := s.UserSpaceRepository.IsUserInSpace(userID, spaceID)
	if err != nil {
		return false, err
	}
	return isInSpace, nil
}
