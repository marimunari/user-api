package middlewares

import (
	"strings"
	"user-api/auth"
	"user-api/utils"

	"github.com/gofiber/fiber/v2"
)

// Middleware for token authentication
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(utils.ErrTokenNotProvided.Code).JSON(utils.ErrTokenNotProvided)
		}

		token = strings.TrimPrefix(token, "Bearer ")

		if auth.IsTokenBlacklisted(token) {
			return c.Status(utils.ErrTokenBlacklisted.Code).JSON(utils.ErrTokenBlacklisted)
		}

		claims, err := auth.ValidateToken(token, false)
		if err != nil {
			return c.Status(utils.ErrInvalidToken.Code).JSON(utils.ErrInvalidToken)
		}

		c.Locals("claims", claims)
		return c.Next()
	}
}
