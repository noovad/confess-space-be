package controller

import (
	"errors"
	"fmt"
	"go_confess_space-project/api/service"
	"go_confess_space-project/dto"
	"go_confess_space-project/helper"
	"go_confess_space-project/helper/responsejson"

	"github.com/gin-gonic/gin"
)

type SpaceController struct {
	spaceService service.SpaceService
}

func NewSpaceAuthController(userService service.SpaceService) *SpaceController {
	return &SpaceController{
		spaceService: userService,
	}
}

func (c *SpaceController) CreateSpace(ctx *gin.Context) {
	var requestBody dto.CreateSpaceRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		responsejson.BadRequest(ctx, err)
		return
	}

	space, err := c.spaceService.CreateSpace(requestBody)
	if err != nil {
		responsejson.InternalServerError(ctx, err)
		return
	}

	responsejson.Created(ctx, space, "Space created successfully")
}

func (c *SpaceController) GetSpaces(ctx *gin.Context) {
	limit := 10
	page := 1
	search := ctx.Query("search")

	if limitStr := ctx.Query("limit"); limitStr != "" {
		if _, err := fmt.Sscanf(limitStr, "%d", &limit); err != nil {
			responsejson.BadRequest(ctx, errors.New("invalid limit"))
			return
		}
	}

	if pageStr := ctx.Query("page"); pageStr != "" {
		if _, err := fmt.Sscanf(pageStr, "%d", &page); err != nil {
			responsejson.BadRequest(ctx, errors.New("invalid page"))
			return
		}
	}

	spaces, err := c.spaceService.GetSpaces(limit, page, search)
	if err != nil {
		responsejson.InternalServerError(ctx, err)
		return
	}

	responsejson.Success(ctx, spaces, "Spaces retrieved successfully")
}

func (c *SpaceController) GetSpaceById(ctx *gin.Context) {
	spaceId := ctx.Param("id")
	if spaceId == "" {
		responsejson.BadRequest(ctx, errors.New("space ID is required"))
		return
	}

	id, err := helper.StringToUUID(spaceId)
	if err != nil {
		responsejson.BadRequest(ctx, fmt.Errorf("invalid space ID format: %w", err))
		return
	}

	space, err := c.spaceService.GetSpaceById(id)
	if err != nil {
		responsejson.InternalServerError(ctx, err)
		return
	}

	responsejson.Success(ctx, space, "Space retrieved successfully")
}

func (c *SpaceController) UpdateSpace(ctx *gin.Context) {
	spaceId := ctx.Param("id")
	if spaceId == "" {
		responsejson.BadRequest(ctx, errors.New("space ID is required"))
		return
	}

	var requestBody dto.UpdateSpaceRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		responsejson.BadRequest(ctx, err)
		return
	}

	id, err := helper.StringToUUID(spaceId)
	if err != nil {
		responsejson.BadRequest(ctx, fmt.Errorf("invalid space ID format: %w", err))
		return
	}

	requestBody.Id = id

	updatedSpace, err := c.spaceService.UpdateSpace(requestBody)
	if err != nil {
		responsejson.InternalServerError(ctx, err)
		return
	}

	responsejson.Success(ctx, updatedSpace, "Space updated successfully")
}

func (c *SpaceController) DeleteSpace(ctx *gin.Context) {
	spaceId := ctx.Param("id")
	if spaceId == "" {
		responsejson.BadRequest(ctx, errors.New("space ID is required"))
		return
	}

	id, err := helper.StringToUUID(spaceId)
	if err != nil {
		responsejson.BadRequest(ctx, fmt.Errorf("invalid space ID format: %w", err))
		return
	}

	err = c.spaceService.DeleteSpace(id)
	if err != nil {
		responsejson.InternalServerError(ctx, err)
		return
	}

	responsejson.Success(ctx, nil, "Space deleted successfully")
}
