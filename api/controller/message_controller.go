package controller

import (
	"errors"
	"go_confess_space-project/api/service"
	"go_confess_space-project/dto"
	customerror "go_confess_space-project/helper/customerrors"
	"go_confess_space-project/helper/responsejson"

	"github.com/gin-gonic/gin"
)

type MessageController struct {
	messageService service.MessageService
}

func NewMessageController(messageService service.MessageService) *MessageController {
	return &MessageController{
		messageService: messageService,
	}
}
func (c *MessageController) CreateMessage(ctx *gin.Context) {
	var requestBody dto.MessageRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		responsejson.BadRequest(ctx, err)
		return
	}

	message, err := c.messageService.CreateMessage(requestBody)
	if err != nil {
		if errors.Is(err, customerror.ErrValidation) {
			responsejson.BadRequest(ctx, err, "Validation error")
			return
		}
		responsejson.InternalServerError(ctx, err)
		return
	}

	responsejson.Created(ctx, message, "Message created successfully")
}

func (c *MessageController) GetMessages(ctx *gin.Context) {
	spaceID := ctx.Param("spaceID")
	if spaceID == "" {
		responsejson.BadRequest(ctx, errors.New("space ID is required"))
		return
	}

	messages, err := c.messageService.GetMessages(spaceID)
	if err != nil {
		responsejson.InternalServerError(ctx, err)
		return
	}

	responsejson.Success(ctx, messages, "Messages retrieved successfully")
}
