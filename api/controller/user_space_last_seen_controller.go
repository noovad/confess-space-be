package controller

import (
	"errors"
	"fmt"
	"go_confess_space-project/api/service"
	"go_confess_space-project/dto"
	customerror "go_confess_space-project/helper/customerrors"
	"go_confess_space-project/helper/responsejson"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responsejson.NotFound(ctx, fmt.Sprintf("Last seen for user %s in space %s not found", userID, spaceID))
			return
		}
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
		if errors.Is(err, customerror.ErrValidation) {
			responsejson.BadRequest(ctx, err, "Validation error")
			return
		}
		if errors.Is(err, customerror.ErrForeignKeyViolation) {
			responsejson.BadRequest(ctx, err, "Foreign key violation")
			return
		}
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responsejson.NotFound(ctx, fmt.Sprintf("Last seen for user %s in space %s not found", userID, spaceID))
			return
		}
		responsejson.InternalServerError(ctx, err)
		return
	}

	responsejson.Success(ctx, nil, "Last seen deleted successfully")
}
