package controller

import (
	"errors"
	"go_confess_space-project/api/service"
	"go_confess_space-project/dto"
	"go_confess_space-project/helper"
	customerror "go_confess_space-project/helper/customerrors"
	"go_confess_space-project/helper/responsejson"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserSpaceController struct {
	userSpaceService service.UserSpaceService
}

func NewUserSpaceController(userSpaceService service.UserSpaceService) *UserSpaceController {
	return &UserSpaceController{
		userSpaceService: userSpaceService,
	}
}

func (c *UserSpaceController) AddUserToSpace(ctx *gin.Context) {
	var req dto.UserSpaceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responsejson.BadRequest(ctx, err, "Invalid request body")
		return
	}

	userSpaceResponse, err := c.userSpaceService.AddUserToSpace(req)
	if err != nil {
		if errors.Is(err, customerror.ErrValidation) {
			responsejson.BadRequest(ctx, err, "Validation error")
			return
		}
		if errors.Is(err, customerror.ErrForeignKeyViolation) {
			responsejson.BadRequest(ctx, err, "Foreign key violation")
			return
		}
		responsejson.InternalServerError(ctx, err, "Failed to add user to space")
		return
	}

	responsejson.Created(ctx, userSpaceResponse, "User added to space successfully")
}

func (c *UserSpaceController) RemoveUserFromSpace(ctx *gin.Context) {
	spaceID := uuid.MustParse(ctx.Param("spaceID"))

	userID, valid := helper.ValidateAccessToken(helper.AccessTokenFromHeader(ctx))
	if !valid {
		responsejson.Unauthorized(ctx, "Invalid access token")
		return
	}

	err := c.userSpaceService.RemoveUserFromSpace(spaceID, uuid.MustParse(userID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responsejson.NotFound(ctx, "User not found in space")
			return
		}
		responsejson.InternalServerError(ctx, err, "Failed to remove user from space")
		return
	}

	responsejson.Success(ctx, nil, "User removed from space successfully")
}

func (c *UserSpaceController) GetUserSpace(ctx *gin.Context) {
	spaceIDParam := ctx.Query("spaceId")
	var spaceID uuid.UUID
	var err error
	if spaceIDParam != "" {
		spaceID, err = uuid.Parse(spaceIDParam)
		if err != nil {
			responsejson.BadRequest(ctx, err, "Invalid spaceId")
			return
		}
	}

	userIDParam := ctx.Query("userId")
	var userID uuid.UUID
	if userIDParam != "" {
		userID, err = uuid.Parse(userIDParam)
		if err != nil {
			responsejson.BadRequest(ctx, err, "Invalid userId")
			return
		}

	}

	userSpaces, err := c.userSpaceService.GetUserSpace(spaceID, userID)
	if err != nil {
		responsejson.InternalServerError(ctx, err, "Failed to retrieve user space")
		return
	}

	responsejson.Success(ctx, userSpaces, "User spaces fetched successfully")
}

func (c *UserSpaceController) IsUserInSpace(ctx *gin.Context) {
	spaceID := uuid.MustParse(ctx.Param("spaceID"))
	userID := uuid.MustParse(ctx.Param("userID"))

	isInSpace, err := c.userSpaceService.IsUserInSpace(spaceID, userID)
	if err != nil {
		responsejson.InternalServerError(ctx, err, "Failed to check if user is in space")
		return
	}

	responsejson.Success(ctx, isInSpace, "User in space check completed")
}
