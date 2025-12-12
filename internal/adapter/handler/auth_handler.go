package handler

import (
	"bwanews/internal/adapter/handler/request"
	"bwanews/internal/adapter/handler/response"
	"bwanews/internal/core/domain/entity"
	"bwanews/internal/core/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

var err error
var code string
var errorResp response.ErrorResponseDefault
var validate = validator.New()

type AuthHandler interface {
	Login(c *fiber.Ctx) error
}

type authHandler struct {
	authService service.AuthService
}

func (a *authHandler)Login(c *fiber.Ctx) error {
	 req := request.LoginRequest{}

	 response := response.SuccessAuthResponse{}

	 if err = c.BodyParser(&req); err != nil {
		code = "[HANDLER Login - 1]"
		log.Errorw(code, err)

		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusBadRequest).JSON(errorResp)
	} 

	if validate.Struct(req); err != nil {
		code = "[HANDLER] Login -2"
		log.Errorw(code, err)

		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusBadRequest).JSON(errorResp)
	}


	reqLogin := entity.LoginRequest{
		Email: req.Email,
		Password: req.Password,
	}

	result, err := a.authService.GetUserByEmail(c.Context(), reqLogin)

	if err != nil {
		code = "[HANDLER] Login -3"
		log.Errorw(code, err)

		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusInternalServerError).JSON(errorResp)
	}

	response.Meta.Status = true
	response.Meta.Message = "Login successfully"
	response.AccessToken = result.AccessToken
	response.ExpiresAt = result.ExpiresAt

	return c.JSON(response)
}

func NewAuthHandler(authService service.AuthService) AuthHandler {
	return &authHandler{
		authService: authService,
	}
}

