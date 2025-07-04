package controller

import (
	"errors"
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
			responsejson.NotFound(ctx, "Last seen Not found")
			return
		}
		responsejson.InternalServerError(ctx, err, "Failed to retrieve last seen")
		return
	}

	responsejson.Success(ctx, lastSeen, "Last seen retrieved successfully")
}

func (c *UserSpaceLastSeenController) CreateOrUpdateLastSeen(ctx *gin.Context) {
	var requestBody dto.UserSpaceLastSeenRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		responsejson.BadRequest(ctx, err, "Invalid request body")
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
		responsejson.InternalServerError(ctx, err, "Failed to create or update last seen")
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
			responsejson.NotFound(ctx, "Last seen Not found")
			return
		}
		responsejson.InternalServerError(ctx, err, "Failed to delete last seen")
		return
	}

	responsejson.Success(ctx, nil, "Last seen deleted successfully")
}
