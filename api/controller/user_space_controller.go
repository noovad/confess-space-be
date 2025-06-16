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
	var requestBody dto.UserSpaceRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		responsejson.BadRequest(ctx, err)
		return
	}

	userSpaceResponse, err := c.userSpaceService.AddUserToSpace(requestBody)
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

	responsejson.Created(ctx, userSpaceResponse, "User added to space successfully")
}

func (c *UserSpaceController) RemoveUserFromSpace(ctx *gin.Context) {
	spaceIDStr := ctx.Param("spaceID")
	spaceID, err := helper.StringToUUID(spaceIDStr)
	if err != nil {
		responsejson.BadRequest(ctx, err)
		return
	}

	userIDStr := ctx.Param("userID")
	userID, err := helper.StringToUUID(userIDStr)
	if err != nil {
		responsejson.BadRequest(ctx, err)
		return
	}

	err = c.userSpaceService.RemoveUserFromSpace(spaceID, userID)
	if err != nil {
		responsejson.InternalServerError(ctx, err)
		return
	}

	responsejson.Success(ctx, nil, "User removed from space successfully")
}

func (c *UserSpaceController) GetUserSpace(ctx *gin.Context) {
	spaceIDStr := ctx.Query("spaceId")
	var spaceID uuid.UUID
	var err error
	if spaceIDStr != "" {
		spaceID, err = helper.StringToUUID(spaceIDStr)
		if err != nil {
			responsejson.BadRequest(ctx, err)
			return
		}
	}

	userIDStr := ctx.Query("userId")
	var userID uuid.UUID
	if userIDStr != "" {
		userID, err = helper.StringToUUID(userIDStr)
		if err != nil {
			responsejson.BadRequest(ctx, err)
			return
		}
	}

	userSpaces, err := c.userSpaceService.GetUserSpace(spaceID, userID)
	if err != nil {
		responsejson.InternalServerError(ctx, err)
		return
	}

	responsejson.Success(ctx, userSpaces, "User spaces fetched successfully")
}

func (c *UserSpaceController) IsUserInSpace(ctx *gin.Context) {
	userIDStr := ctx.Param("userID")
	spaceIDStr := ctx.Param("spaceID")

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		responsejson.BadRequest(ctx, err)
		return
	}

	spaceID, err := uuid.Parse(spaceIDStr)
	if err != nil {
		responsejson.BadRequest(ctx, err)
		return
	}

	isInSpace, err := c.userSpaceService.IsUserInSpace(userID, spaceID)
	if err != nil {
		responsejson.InternalServerError(ctx, err)
		return
	}

	responsejson.Success(ctx, isInSpace, "User in space check completed")
}
