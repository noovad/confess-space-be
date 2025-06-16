package service

import (
	"go_confess_space-project/api/repository"
	"go_confess_space-project/dto"
	"go_confess_space-project/model"

	"github.com/go-playground/validator/v10"
)

type UserSpaceLastSeenService interface {
	GetLastSeenByUserAndSpace(userID string, spaceID string) (model.UserSpaceLastSeen, error)
	CreateOrUpdateLastSeen(request dto.UserSpaceLastSeenRequest) (model.UserSpaceLastSeen, error)
	DeleteLastSeenByUserAndSpace(userID string, spaceID string) error
}

type UserSpaceLastSeenServiceImpl struct {
	Repository repository.UserSpaceLastSeenRepository
	Validate   *validator.Validate
}

func NewUserSpaceLastSeenServiceImpl(repository repository.UserSpaceLastSeenRepository, validate *validator.Validate) UserSpaceLastSeenService {
	return &UserSpaceLastSeenServiceImpl{
		Repository: repository,
		Validate:   validate,
	}
}

func (s *UserSpaceLastSeenServiceImpl) GetLastSeenByUserAndSpace(userID string, spaceID string) (model.UserSpaceLastSeen, error) {
	lastSeen, err := s.Repository.GetLastSeenByUserAndSpace(userID, spaceID)
	if err != nil {
		return model.UserSpaceLastSeen{}, err
	}
	return lastSeen, nil
}

func (s *UserSpaceLastSeenServiceImpl) CreateOrUpdateLastSeen(request dto.UserSpaceLastSeenRequest) (model.UserSpaceLastSeen, error) {
	if err := s.Validate.Struct(request); err != nil {
		return model.UserSpaceLastSeen{}, err
	}

	requestModel := model.UserSpaceLastSeen{
		UserID:   request.UserID,
		SpaceID:  request.SpaceID,
		LastSeen: request.LastSeen,
	}

	result, err := s.Repository.CreateOrUpdateLastSeen(requestModel)
	if err != nil {
		return model.UserSpaceLastSeen{}, err
	}
	return result, nil
}

func (s *UserSpaceLastSeenServiceImpl) DeleteLastSeenByUserAndSpace(userID string, spaceID string) error {
	err := s.Repository.DeleteLastSeenByUserAndSpace(userID, spaceID)
	if err != nil {
		return err
	}
	return nil
}
