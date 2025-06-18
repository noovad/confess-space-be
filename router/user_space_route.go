package router

import (
	"go_confess_space-project/api"
	"go_confess_space-project/helper"

	"github.com/gin-gonic/gin"
)

func UserSpaceRoutes(r *gin.RouterGroup) {
	authMidleware := helper.AuthMiddleware
	UserSpaceController := api.UserSpaceInjector()

	{
		space := r.Group("/user-space")
		space.Use(authMidleware)
		space.POST("", UserSpaceController.AddUserToSpace)
		space.DELETE("/:spaceID/:userID", UserSpaceController.RemoveUserFromSpace)
		space.GET("", UserSpaceController.GetUserSpace)
		space.GET("/check/:spaceID/:userID", UserSpaceController.IsUserInSpace)
	}
}
