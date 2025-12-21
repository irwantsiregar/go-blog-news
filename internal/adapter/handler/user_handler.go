package handler

import (
	"bwanews/internal/adapter/handler/request"
	"bwanews/internal/adapter/handler/response"
	"bwanews/internal/core/domain/entity"
	"bwanews/internal/core/service"
	validatorLib "bwanews/lib/validator"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type UserHandler interface {
	UpdatePassword(c *fiber.Ctx) error
	GetUserByID(c *fiber.Ctx) error
}

type userHandler struct {
	userService service.UserService
}

// GetUserByID implements UserHandler.
func (u *userHandler) GetUserByID(c *fiber.Ctx) error {
	claims := c.Locals("user").(*entity.JwtData)

	userID := claims.UserID

	if userID == 0 {
		code = "[HANDLER] GetUserByID - 1"
		log.Errorw(code, err)

		errorResp.Meta.Status = false
		errorResp.Meta.Message = "Unauthorized access"
		return c.Status(fiber.StatusUnauthorized).JSON(errorResp)
	}

	user, err := u.userService.GetUserByID(c.Context(), int64(userID))

	if err != nil {
		code = "[HANDLER] GetUserByID - 2"
		log.Errorw(code, err)

		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusInternalServerError).JSON(errorResp)
	}


	response := response.UserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
	}
	
	defaultSuccessResponse.Meta.Status = true
	defaultSuccessResponse.Meta.Message = "Success"
	defaultSuccessResponse.Data = response

	return c.JSON(defaultSuccessResponse)
}

// UpdatePassword implements UserHandler.
func (u *userHandler) UpdatePassword(c *fiber.Ctx) error {
	claims := c.Locals("user").(*entity.JwtData)

	userID := claims.UserID

	if userID == 0 {
		code = "[HANDLER] UpdatePassword - 1"
		log.Errorw(code, err)

		errorResp.Meta.Status = false
		errorResp.Meta.Message = "Unauthorized access"
		return c.Status(fiber.StatusUnauthorized).JSON(errorResp)
	}

	
	var req request.UpdatePasswordRequest

	if err = c.BodyParser(&req); err != nil {
		code = "[HANDLER] UpdatePassword - 2"
		log.Errorw(code, err)

		errorResp.Meta.Status = false
		errorResp.Meta.Message = "Invalid request body"  
		return c.Status(fiber.StatusInternalServerError).JSON(errorResp)
	}

	if err = validatorLib.ValidateStruct(req); err != nil {
		code = "[HANDLER] UpdatePassword - 3"
		err = errors.New("Password does not match")
		log.Errorw(code, err)
		
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusBadRequest).JSON(errorResp)
	}

	if req.NewPassword != req.ConfirmPassword {
		code = "[HANDLER] UpdatePassword - 4"

		log.Errorw(code, err)

		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusBadRequest).JSON(errorResp)
	}

	err = u.userService.UpdatePassword(c.Context(), req.NewPassword, int64(userID))

	if err != nil {
		code = "[HANDLER] UpdatePassword - 5"
		log.Errorw(code, err)

		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusInternalServerError).JSON(errorResp)
	}

	
	defaultSuccessResponse.Meta.Status = true
	defaultSuccessResponse.Meta.Message = "Success"
	defaultSuccessResponse.Data = nil

	return c.JSON(defaultSuccessResponse)

}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{userService: userService}
}
