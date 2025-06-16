package controller

import (
	"go_confess_space-project/api/service"
	"go_confess_space-project/dto"
	"go_confess_space-project/helper/responsejson"

	"github.com/gin-gonic/gin"
)

type UserSpaceLastSeenController struct {
	UserSpaceLastSeenService service.UserSpaceLastSeenService
}

func NewUserSpaceLastSeenController(userSpaceLastSeenService service.UserSpaceLastSeenService) *UserSpaceLastSeenController {
	return &UserSpaceLastSeenController{
		UserSpaceLastSeenService: userSpaceLastSeenService,
	}
}
func (c *UserSpaceLastSeenController) GetLastSeen(ctx *gin.Context) {
	userID := ctx.Param("userID")
	spaceID := ctx.Param("spaceID")

	lastSeen, err := c.UserSpaceLastSeenService.GetLastSeenByUserAndSpace(userID, spaceID)
	if err != nil {
		responsejson.InternalServerError(ctx, err)
		return
	}

	responsejson.Success(ctx, lastSeen, "Last seen retrieved successfully")
}

func (c *UserSpaceLastSeenController) CreateOrUpdateLastSeen(ctx *gin.Context) {
	var requestBody dto.UserSpaceLastSeenRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		responsejson.BadRequest(ctx, err)
		return
	}

	lastSeen, err := c.UserSpaceLastSeenService.CreateOrUpdateLastSeen(requestBody)
	if err != nil {
		responsejson.InternalServerError(ctx, err)
		return
	}

	responsejson.Success(ctx, lastSeen, "Last seen created or updated successfully")
}

func (c *UserSpaceLastSeenController) DeleteLastSeen(ctx *gin.Context) {
	userID := ctx.Param("userID")
	spaceID := ctx.Param("spaceID")

	err := c.UserSpaceLastSeenService.DeleteLastSeenByUserAndSpace(userID, spaceID)
	if err != nil {
		responsejson.InternalServerError(ctx, err)
		return
	}

	responsejson.Success(ctx, nil, "Last seen deleted successfully")
}
