package service

import (
	"go_confess_space-project/api/repository"
	"go_confess_space-project/dto"
	customerror "go_confess_space-project/helper/customerrors"
	"go_confess_space-project/model"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type MessageService interface {
	CreateMessage(message dto.MessageRequest, id string) (model.Message, error)
	GetMessages(spaceID string) ([]model.Message, error)
}

func NewMessageServiceImpl(messageRepository repository.MessageRepository,
	validate *validator.Validate,
) MessageService {
	return &MessageServiceImpl{
		MessageRepository: messageRepository,
		Validate:          validate,
	}
}

type MessageServiceImpl struct {
	MessageRepository repository.MessageRepository
	Validate          *validator.Validate
}

func (s *MessageServiceImpl) CreateMessage(req dto.MessageRequest, id string) (model.Message, error) {
	err := s.Validate.Struct(req)

	if err != nil {
		return model.Message{}, customerror.WrapValidation(err)
	}

	message := model.Message{
		SpaceID: uuid.MustParse(req.SpaceID),
		UserID:  uuid.MustParse(id),
		Content: req.Message,
	}

	createdMessage, err := s.MessageRepository.CreateMessage(message)
	if err != nil {
		return model.Message{}, customerror.HandlePostgresError(err)
	}
	return createdMessage, nil
}

func (s *MessageServiceImpl) GetMessages(spaceID string) ([]model.Message, error) {
	return s.MessageRepository.GetMessages(spaceID)
}
