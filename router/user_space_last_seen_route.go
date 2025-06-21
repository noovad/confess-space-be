package router

import (
	"go_confess_space-project/api"

	"github.com/gin-gonic/gin"
	"github.com/noovad/go-auth/helper"
)

func UserSpaceLastSeenRoutes(r *gin.RouterGroup) {
	authMiddleware := helper.AuthMiddleware
	UserSpaceLastSeenController := api.UserSpaceLastSeenInjector()

	{
		userSpaceLastSeen := r.Group("/user-space-last-seen")
		userSpaceLastSeen.Use(authMiddleware)
		userSpaceLastSeen.GET("/:spaceID/:userID", UserSpaceLastSeenController.GetLastSeen)
		userSpaceLastSeen.POST("", UserSpaceLastSeenController.CreateOrUpdateLastSeen)
		userSpaceLastSeen.DELETE("/:spaceID/:userID", UserSpaceLastSeenController.DeleteLastSeen)
	}
}
