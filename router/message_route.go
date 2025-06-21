package router

import (
	"go_confess_space-project/api"

	"github.com/gin-gonic/gin"
	"github.com/noovad/go-auth/helper"
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
