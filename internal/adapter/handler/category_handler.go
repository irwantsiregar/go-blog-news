package handler

import (
	"bwanews/internal/adapter/handler/response"
	"bwanews/internal/core/domain/entity"
	"bwanews/internal/core/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

var defaultSuccessResponse response.DefaultSuccessResponse

type CategoryHandler interface {
	GetCategories(c *fiber.Ctx) 
	GetCategoryByID(c *fiber.Ctx)
	CreateCategory(c *fiber.Ctx) 
	EditCategoryByID(c *fiber.Ctx)
	DeleteCategory(c *fiber.Ctx) 
}

type categoryHandler struct {
	categoryService service.CategoryService
}

// CreateCategory implements CategorySHandler.
func (ch *categoryHandler) CreateCategory(c *fiber.Ctx) error {
	panic("unimplemented")
}

// DeleteCategory implements CategorySHandler.
func (ch *categoryHandler) DeleteCategory(c *fiber.Ctx) error {
	panic("unimplemented")
}

// EditCategoryByID implements CategorySHandler.
func (ch *categoryHandler) EditCategoryByID(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetCategories implements CategorySHandler.
func (ch *categoryHandler) GetCategories(c *fiber.Ctx) (error) {
	claims := c.Locals("user").(*entity.JwtData)

	userID := claims.UserID

	if userID == 0 {
		code = "[HANDLER] GetCategories - 1"
		log.Errorw(code, err)

		errorResp.Meta.Status = false
		errorResp.Meta.Message = "Unauthorized access"
		return c.Status(fiber.StatusUnauthorized).JSON(errorResp)
	}

	results, err := ch.categoryService.GetCategories(c.Context())

	if err != nil {
		code = "[HANDLER] GetCategories - 2"
		log.Errorw(code, err)

		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(errorResp)
	}

	categoryResponses := []response.SuccessCategoryResponse{}

	for _, result := range results {
		categoryResponse := response.SuccessCategoryResponse{
			ID: result.ID,
			Title : result.Title,
			Slug : result.Slug,
			CreatedByName : result.User.Name,
		}

		categoryResponses = append(categoryResponses, categoryResponse)
	}

	defaultSuccessResponse.Meta.Status = true
	defaultSuccessResponse.Meta.Message = "Categories fetched succesfully"
	defaultSuccessResponse.Data = categoryResponses

	return c.JSON(defaultSuccessResponse)
}

// GetCategoryByID implements CategorySHandler.
func (ch *categoryHandler) GetCategoryByID(c *fiber.Ctx, id int64) ([]entity.CategoryEntity, error) {
	panic("unimplemented")
}


func NewCategoryHandler(categoryService service.CategoryService) *categoryHandler {
	return &categoryHandler{
		categoryService: categoryService,
	}
}