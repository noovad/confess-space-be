package service

import (
	"go_confess_space-project/api/repository"
	"go_confess_space-project/dto"
	customerror "go_confess_space-project/helper/customerrors"
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
	return s.Repository.GetLastSeenByUserAndSpace(userID, spaceID)
}

func (s *UserSpaceLastSeenServiceImpl) CreateOrUpdateLastSeen(req dto.UserSpaceLastSeenRequest) (model.UserSpaceLastSeen, error) {
	if err := s.Validate.Struct(req); err != nil {
		return model.UserSpaceLastSeen{}, customerror.WrapValidation(err)
	}

	reqModel := model.UserSpaceLastSeen{
		UserID:   req.UserID,
		SpaceID:  req.SpaceID,
		LastSeen: req.LastSeen,
	}

	return s.Repository.CreateOrUpdateLastSeen(reqModel)
}

func (s *UserSpaceLastSeenServiceImpl) DeleteLastSeenByUserAndSpace(userID string, spaceID string) error {
	return s.Repository.DeleteLastSeenByUserAndSpace(userID, spaceID)
}
