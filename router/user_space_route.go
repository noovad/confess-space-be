package router

import (
	"go_confess_space-project/api"

	"github.com/gin-gonic/gin"
	"github.com/noovad/go-auth/helper"
)

func UserSpaceRoutes(r *gin.RouterGroup) {
	authMiddleware := helper.AuthMiddleware
	UserSpaceController := api.UserSpaceInjector()

	{
		space := r.Group("/user-space")
		space.Use(authMiddleware)
		space.POST("", UserSpaceController.AddUserToSpace)
		space.DELETE("/:spaceId", UserSpaceController.RemoveUserFromSpace)
		space.GET("", UserSpaceController.GetUserSpace)
		space.GET("/check/:spaceID/:userID", UserSpaceController.IsUserInSpace)
	}
}
