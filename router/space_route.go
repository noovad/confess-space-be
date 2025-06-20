package router

import (
	"go_confess_space-project/api"
	"go_confess_space-project/helper"

	"github.com/gin-gonic/gin"
)

func SpaceRoutes(r *gin.RouterGroup) {
	authMidleware := helper.AuthMiddleware
	SpaceController := api.SpaceInjector()

	{
		space := r.Group("/space")
		space.Use(authMidleware)
		space.POST("", SpaceController.CreateSpace)
		space.GET("", SpaceController.GetSpaces)
		space.GET("/:id", SpaceController.GetSpaceById)
		space.GET("/slug/:slug", SpaceController.GetSpaceBySlug)
		space.PUT("/:id", SpaceController.UpdateSpace)
		space.DELETE("/:id", SpaceController.DeleteSpace)
	}
}
