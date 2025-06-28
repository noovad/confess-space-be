package router

import (
	"go_confess_space-project/api"
	"go_confess_space-project/api/controller"

	"github.com/gin-gonic/gin"
	"github.com/noovad/go-auth/helper"
)

func MessageRoutes(r *gin.RouterGroup) {
	authMiddleware := helper.AuthMiddleware
	wsController := controller.NewWebSocketController()
	messageController := api.MessageInjector()

	{
		r.GET("/ws/connect", wsController.HandleWebSocket)

		r.POST("/messages", authMiddleware, messageController.CreateMessage)
		r.GET("/messages/channel/:channelID", authMiddleware, messageController.GetChannelMessages)
	}
}
