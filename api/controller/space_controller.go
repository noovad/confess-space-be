package controller

import (
	"errors"
	"fmt"
	"go_confess_space-project/api/service"
	"go_confess_space-project/dto"
	customerror "go_confess_space-project/helper/customerrors"
	"go_confess_space-project/helper/responsejson"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
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
		responsejson.BadRequest(ctx, err, "Invalid request body")
		return
	}

	userId, exists := ctx.Get("userId")
	if !exists {
		responsejson.InternalServerError(ctx, nil, "User ID not found in context")
		return
	}

	space, err := c.spaceService.CreateSpace(requestBody, userId.(string))
	if err != nil {
		if errors.Is(err, customerror.ErrValidation) {
			responsejson.BadRequest(ctx, err, "Validation error")
			return
		}
		if errors.Is(err, customerror.ErrUniqueViolation) {
			responsejson.Conflict(ctx, err, "Unique constraint violation")
			return
		}
		if errors.Is(err, customerror.ErrForeignKeyViolation) {
			responsejson.BadRequest(ctx, err, "Foreign key violation")
			return
		}
		responsejson.InternalServerError(ctx, err, "Failed to create space")
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
			responsejson.BadRequest(ctx, err, "invalid limit")
			return
		}
	}

	if pageStr := ctx.Query("page"); pageStr != "" {
		if _, err := fmt.Sscanf(pageStr, "%d", &page); err != nil {
			responsejson.BadRequest(ctx, err, "invalid page")
			return
		}
	}

	spaces, err := c.spaceService.GetSpaces(limit, page, search)
	if err != nil {
		responsejson.InternalServerError(ctx, err, "Failed to retrieve spaces")
		return
	}

	responsejson.Success(ctx, spaces, "Spaces retrieved successfully")
}

func (c *SpaceController) GetSpaceById(ctx *gin.Context) {
	spaceId := ctx.Param("id")

	id := uuid.MustParse(spaceId)

	space, err := c.spaceService.GetSpaceById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responsejson.NotFound(ctx, "Space Not Found")
			return
		}
		responsejson.InternalServerError(ctx, err, "Failed to retrieve space")
		return
	}

	responsejson.Success(ctx, space, "Space retrieved successfully")
}

func (c *SpaceController) GetSpaceBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")
	if slug == "" {
		responsejson.BadRequest(ctx, errors.New("slug is required"))
		return
	}

	space, err := c.spaceService.GetSpaceBySlug(slug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responsejson.NotFound(ctx, "Space Not Found")
			return
		}
		responsejson.InternalServerError(ctx, err, "Failed to retrieve space")
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
		responsejson.BadRequest(ctx, err, "Invalid request body")
		return
	}

	requestBody.Id = uuid.MustParse(spaceId)

	updatedSpace, err := c.spaceService.UpdateSpace(requestBody)
	if err != nil {
		if errors.Is(err, customerror.ErrValidation) {
			responsejson.BadRequest(ctx, err, "Validation error")
			return
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responsejson.NotFound(ctx, "Space Not Found")
			return
		}
		responsejson.InternalServerError(ctx, err, "Failed to update space")
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

	id := uuid.MustParse(spaceId)

	err := c.spaceService.DeleteSpace(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responsejson.NotFound(ctx, "Space Not Found")
			return
		}
		responsejson.InternalServerError(ctx, err, "Failed to delete space")
		return
	}

	responsejson.Success(ctx, nil, "Space deleted successfully")
}
