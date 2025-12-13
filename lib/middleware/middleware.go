package middleware

import (
	"bwanews/config"
	"bwanews/internal/adapter/handler/response"
	"bwanews/lib/auth"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Middleware interface {
	CheckToken() fiber.Handler
}

type Options struct {
	authJwt auth.Jwt
}

func (o *Options) CheckToken() fiber.Handler { 
	return func(c *fiber.Ctx) error {
		var errorResponse response.ErrorResponseDefault

		authHeader := c.Get("Authorization")	
		
		if authHeader == "" {
			errorResponse.Meta.Status = false
			errorResponse.Meta.Message = "Missing Authorization Header"

			return c.Status(fiber.StatusUnauthorized).JSON(errorResponse)
		}		

		// tokenString := authHeader[len("Bearer "):]
		tokenString := strings.Split(authHeader, "Bearer ")[1]

		jwtData, err := o.authJwt.VerifyAccessToken(tokenString)
		
		if err != nil {
			errorResponse.Meta.Status = false
			errorResponse.Meta.Message = "Invalid or expired JWT"

			return c.Status(fiber.StatusUnauthorized).JSON(errorResponse)
		}
		
		c.Locals("user", jwtData)
		
		return c.Next()
	}	
}

func NewMiddleware(cfg *config.Config) Middleware {
	opt := new(Options)
	
	opt.authJwt = auth.NewJwt(cfg)

	return opt
}














