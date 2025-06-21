package router

import (
	"go_confess_space-project/api"

	"github.com/gin-gonic/gin"
	"github.com/noovad/go-auth/helper"
)

func SpaceRoutes(r *gin.RouterGroup) {
	authMiddleware := helper.AuthMiddleware
	SpaceController := api.SpaceInjector()

	{
		space := r.Group("/space")
		space.Use(authMiddleware)
		space.POST("", SpaceController.CreateSpace)
		space.GET("/own", SpaceController.GetOwnSpace)
		space.GET("", SpaceController.GetSpaces)
		space.GET("/slug/:slug", SpaceController.GetSpaceBySlug)
		space.PUT("/:id", SpaceController.UpdateSpace)
		space.DELETE("/:id", SpaceController.DeleteSpace)
		space.GET("/exists", SpaceController.ExistsByOwnerID)
	}
}
