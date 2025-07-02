package router

import (
	"go_confess_space-project/api"
	"go_confess_space-project/api/controller"
	"go_confess_space-project/config/websocket"

	"github.com/gin-gonic/gin"
	"github.com/noovad/go-auth/helper"
)

func MessageRoutes(r *gin.RouterGroup, hub *websocket.Hub) {
	authMiddleware := helper.AuthMiddleware
	wsController := controller.NewWebSocketController(hub)
	messageController := controller.NewMessageController(
		*api.MessageInjector(),
		hub,
	)

	{
		r.GET("/ws/connect", wsController.HandleWebSocket)

		r.POST("/message", authMiddleware, messageController.CreateMessage)
		r.GET("/messages/:channelID", authMiddleware, messageController.GetChannelMessages)
	}
}
