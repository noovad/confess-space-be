package controller

import (
	"errors"
	"go_confess_space-project/api/service"
	"go_confess_space-project/dto"
	"go_confess_space-project/helper"
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
		responsejson.BadRequest(ctx, err, "Invalid request body")
		return
	}

	id, valid := helper.ValidateAccessToken(helper.AccessTokenFromHeader(ctx))
	if !valid {
		responsejson.Unauthorized(ctx, "Invalid access token")
		return
	}

	message, err := c.messageService.CreateMessage(requestBody, id)
	if err != nil {
		if errors.Is(err, customerror.ErrValidation) {
			responsejson.BadRequest(ctx, err, "Validation error")
			return
		}
		if errors.Is(err, customerror.ErrForeignKeyViolation) {
			responsejson.BadRequest(ctx, err, "Foreign key violation")
			return
		}
		responsejson.InternalServerError(ctx, err,  "Failed to create message")
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
		responsejson.InternalServerError(ctx, err, "Failed to retrieve messages")
		return
	}

	responsejson.Success(ctx, messages, "Messages retrieved successfully")
}
