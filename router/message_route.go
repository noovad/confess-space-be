package router

import (
	"go_confess_space-project/api"
	"go_confess_space-project/helper"

	"github.com/gin-gonic/gin"
)

func MessageRoutes(r *gin.RouterGroup) {
	authMiddleware := helper.AuthMiddleware
	MessageController := api.MessageInjector()

	{
		message := r.Group("/message")
		message.Use(authMiddleware)
		message.POST("", MessageController.CreateMessage)
		message.GET("/:spaceID", MessageController.GetMessages)
	}
}
