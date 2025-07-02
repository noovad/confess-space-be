package controller

import (
	"go_confess_space-project/api/service"
	"go_confess_space-project/config/websocket"
	"go_confess_space-project/dto"
	"go_confess_space-project/helper/responsejson"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MessageController struct {
	messageService service.MessageService
	hub            *websocket.Hub
}

func NewMessageController(messageService service.MessageService, hub *websocket.Hub) *MessageController {
	return &MessageController{
		messageService: messageService,
		hub:            hub,
	}
}

func (c *MessageController) CreateMessage(ctx *gin.Context) {
	channel := ctx.Query("channel")
	if channel == "" {
		responsejson.BadRequest(ctx, nil, "Channel parameter is required")
		return
	}
	var requestBody dto.MessageRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		responsejson.BadRequest(ctx, err, "Invalid request body")
		return
	}

	userID, exists := ctx.Get("userId")
	if !exists {
		responsejson.Unauthorized(ctx, "User ID not found in context")
		return
	}

	message, err := c.messageService.CreateMessage(requestBody, userID.(uuid.UUID).String())
	if err != nil {
		responsejson.InternalServerError(ctx, err, "Failed to create message")
		return
	}

	wsMessage := websocket.Message{
		Type:      websocket.MessageTypeChat,
		ID:        message.ID.String(),
		Message:   message,
		Channel:   channel,
		CreatedAt: time.Now(),
	}

	c.hub.Broadcast <- wsMessage

	responsejson.Created(ctx, message, "Message sent successfully")
}

func (c *MessageController) GetChannelMessages(ctx *gin.Context) {
	channelID := ctx.Param("channelID")

	messages, err := c.messageService.GetMessages(channelID)
	if err != nil {
		responsejson.InternalServerError(ctx, err, "Failed to retrieve messages")
		return
	}

	responsejson.Success(ctx, messages, "Messages retrieved successfully")
}
